package hub

import (
	"context"
	"encoding/json"
	"time"

	"go-cover-parroto/internal/core/logger"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 4096
)

type IncomingMessage struct {
	Content string `json:"content"`
}

type MessageHandler func(ctx context.Context, userID string, content string) error

type Client struct {
	hub      *Hub
	conn     *websocket.Conn
	send     chan []byte
	userID   string
	onSend   MessageHandler
	rootCtx  context.Context
}

func NewClient(hub *Hub, conn *websocket.Conn, userID string, onSend MessageHandler, rootCtx context.Context) *Client {
	return &Client{
		hub:     hub,
		conn:    conn,
		send:    make(chan []byte, 64),
		userID:  userID,
		onSend:  onSend,
		rootCtx: rootCtx,
	}
}

func (c *Client) Start() {
	c.hub.Register(c)
	go c.writePump()
	go c.readPump()
}

func (c *Client) readPump() {
	log := logger.S().With("component", "chat-client", "userId", c.userID)
	defer func() {
		c.hub.Unregister(c)
		_ = c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		return c.conn.SetReadDeadline(time.Now().Add(pongWait))
	})

	for {
		_, raw, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
				log.Warnw("unexpected ws close", "error", err)
			}
			return
		}

		var msg IncomingMessage
		if err := json.Unmarshal(raw, &msg); err != nil {
			c.sendError("invalid message format")
			continue
		}

		if err := c.onSend(c.rootCtx, c.userID, msg.Content); err != nil {
			c.sendError(err.Error())
		}
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		_ = c.conn.Close()
	}()

	for {
		select {
		case payload, ok := <-c.send:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				_ = c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.conn.WriteMessage(websocket.TextMessage, payload); err != nil {
				return
			}

		case <-ticker.C:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Client) sendError(reason string) {
	payload, err := json.Marshal(map[string]any{
		"type": "error",
		"data": map[string]string{"message": reason},
	})
	if err != nil {
		return
	}
	select {
	case c.send <- payload:
	default:
	}
}
