package entity

import (
	"github.com/gorilla/websocket"
)

// BroadCaster entity
type BroadCaster struct {
	Clients   map[*websocket.Conn]string
	Broadcast chan ChatMessage
}
