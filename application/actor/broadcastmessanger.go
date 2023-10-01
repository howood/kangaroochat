package actor

import (
	"context"

	"github.com/gorilla/websocket"
	"github.com/howood/kangaroochat/domain/entity"
	"github.com/howood/kangaroochat/domain/repository"
	log "github.com/howood/kangaroochat/infrastructure/logger"
)

// BroadCastMessanger is BroadCastMessanger struct
type BroadCastMessanger struct {
	repository.BroadCasterRepository
}

// NewBroadCastMessanger creates a new BroadCasterRepository
func NewBroadCastMessanger(ctx context.Context) *BroadCastMessanger {
	return &BroadCastMessanger{
		&broadCastMessange{
			broadCaster: &entity.BroadCaster{
				Clients:   make(map[*websocket.Conn]entity.Client),
				Broadcast: make(chan entity.ChatMessage),
			},
			ctx: ctx,
		},
	}
}

// broadCastMessange is broadCastMessange struct
type broadCastMessange struct {
	broadCaster *entity.BroadCaster
	ctx         context.Context
}

// SetNewClient is set new Client
func (bcm *broadCastMessange) SetNewClient(websocket *websocket.Conn, identifier, clientid string) {
	bcm.broadCaster.Clients[websocket] = entity.Client{
		Identifier: identifier,
		ClientID:   clientid,
	}
}

// DeleteClient is delete Client
func (bcm *broadCastMessange) DeleteClient(websocket *websocket.Conn) {
	delete(bcm.broadCaster.Clients, websocket)
}

// ReadMessage is send message
func (bcm *broadCastMessange) ReadMessage(websocket *websocket.Conn, identifier, clientid, username string) error {
	// Read
	var message entity.ChatMessage
	err := websocket.ReadJSON(&message)
	if err != nil {
		bcm.DeleteClient(websocket)
		return err
	}
	message.Identifier = identifier
	message.ClientID = clientid
	message.UserName = username
	bcm.broadCaster.Broadcast <- message
	return nil
}

// BroadcastMessages is send Broadcast Messages
func (bcm *broadCastMessange) BroadcastMessages() {
	for {
		// メッセージ受け取り
		message := <-bcm.broadCaster.Broadcast
		// クライアントの数だけループ
		for clientsocket, client := range bcm.broadCaster.Clients {
			//　書き込む
			if message.Identifier == client.Identifier {
				if message.ClientID == client.ClientID {
					message.IsSelf = true
				}
				err := clientsocket.WriteJSON(message)
				if err != nil {
					log.Error(bcm.ctx, err)
					clientsocket.Close()
					bcm.DeleteClient(clientsocket)
				}
			}
		}
	}
}
