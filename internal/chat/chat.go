package chat

import (
	"github.com/gorilla/websocket"
	"github.com/yalagtyarzh/online-chat/internal/database"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var (
	// Time allowed to write a message to the peer
	writeWait = 10 * time.Second

	// Maximum message size allowed from peer
	maxMessageSize int64 = 512

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	defaultBroadcastQueueSize = 10000
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
		// return r.Header.Get("Origin") != "http://"+r.Host
	},
}

// Chat represents the chat application.
type Chat struct {
	// broadcast sends the message to a room.
	broadcast chan Message

	// quit signals termination of the goroutine that is handling the
	// broadcast.
	quit chan struct{}

	// In-memory data structure, fetch from db if it doesn't exist.
	// Table that maps session id -> session struct. One-to-one.
	sessions *Sessions

	// Table that maps session id -> user id and vice versa. Many-to-one.
	lookup *Table

	// Table that maps user id -> room ids and vice versa. Many-to-many.
	rooms  *TableCache
	db     *database.Conn
	logger *zap.Logger
}
