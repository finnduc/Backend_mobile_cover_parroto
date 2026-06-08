package hub

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/r3labs/sse/v2"
)

const MessagesStream = "messages"

type SSEHub struct {
	server *sse.Server
	done   chan struct{}
}

func NewSSEHub() *SSEHub {
	server := sse.New()

	server.AutoStream = false
	server.AutoReplay = false
	server.BufferSize = 64
	server.SplitData = true

	server.Headers = map[string]string{
		"Cache-Control":     "no-cache, no-transform",
		"X-Accel-Buffering": "no",
	}

	server.CreateStream(MessagesStream)

	h := &SSEHub{
		server: server,
		done:   make(chan struct{}),
	}

	go h.startPing(MessagesStream, 15*time.Second)

	return h
}

func (h *SSEHub) GetServer() *sse.Server {
	return h.server
}

func (h *SSEHub) Publish(streamID string, eventName string, data any) bool {
	payload, err := json.Marshal(data)
	if err != nil {
		return false
	}

	return h.server.TryPublish(streamID, &sse.Event{
		Event: []byte(eventName),
		Data:  payload,
	})
}

func (h *SSEHub) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.server.ServeHTTP(w, r)
}

func (h *SSEHub) startPing(streamID string, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-h.done:
			return
		case t := <-ticker.C:
			payload, _ := json.Marshal(map[string]any{
				"ts": t.Unix(),
			})

			h.server.TryPublish(streamID, &sse.Event{
				Event: []byte("ping"),
				Data:  payload,
			})
		}
	}
}

func (h *SSEHub) Close() {
	close(h.done)
	h.server.Close()
}
