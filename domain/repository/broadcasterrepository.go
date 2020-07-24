package repository

import "github.com/gorilla/websocket"

// BroadCasterRepository interface
type BroadCasterRepository interface {
	SetNewClient(websocket *websocket.Conn, identifier, clientid string)
	DeleteClient(websocket *websocket.Conn)
	ReadMessage(websocket *websocket.Conn, identifier, clientid string) error
	BroadcastMessages()
}
