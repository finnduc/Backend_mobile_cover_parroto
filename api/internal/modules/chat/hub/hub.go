package hub

import (
	"sync"

	"go-cover-parroto/internal/core/logger"
)

type Hub struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
	mu         sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client, 16),
		unregister: make(chan *Client, 16),
		broadcast:  make(chan []byte, 256),
	}
}

func (h *Hub) Run() {
	log := logger.S().With("component", "chat-hub")
	for {
		select {
		case c := <-h.register:
			h.mu.Lock()
			h.clients[c] = true
			count := len(h.clients)
			h.mu.Unlock()
			log.Infow("client registered", "userId", c.userID, "total", count)

		case c := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[c]; ok {
				delete(h.clients, c)
				close(c.send)
			}
			count := len(h.clients)
			h.mu.Unlock()
			log.Infow("client unregistered", "userId", c.userID, "total", count)

		case payload := <-h.broadcast:
			h.mu.RLock()
			for c := range h.clients {
				select {
				case c.send <- payload:
				default:
					// client buffer full → drop client to avoid blocking the hub
					go func(client *Client) { h.unregister <- client }(c)
				}
			}
			h.mu.RUnlock()
		}
	}
}

func (h *Hub) Register(c *Client)         { h.register <- c }
func (h *Hub) Unregister(c *Client)       { h.unregister <- c }
func (h *Hub) Broadcast(payload []byte)   { h.broadcast <- payload }
func (h *Hub) ClientCount() int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(h.clients)
}
