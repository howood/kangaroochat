package actor

import (
	"bytes"
	"context"
	"encoding/gob"

	"github.com/howood/kangaroochat/domain/entity"
	"github.com/howood/kangaroochat/domain/repository"
	"github.com/howood/kangaroochat/infrastructure/uuid"
)

// ChatRoomOperator struct
type ChatRoomOperator struct {
	roomData entity.ChatRoom
	ctx      context.Context
}

// NewChatRoomOperator creates a new ChatRoomRepository
func NewChatRoomOperator(ctx context.Context) repository.ChatRoomRepository {
	return &ChatRoomOperator{ctx: ctx}
}

//Set sets roomname and password to  roomdata
func (e *ChatRoomOperator) Set(roomname, password string) error {
	hashedpassword, salt, err := PasswordOperator{}.GetHashedPassword(password)
	if err != nil {
		return err
	}
	e.roomData.Identifier = uuid.GetUUID(uuid.SegmentioKsuid)
	e.roomData.RoomName = roomname
	e.roomData.HashedPassword = hashedpassword
	e.roomData.Salt = salt
	return nil
}

// GetRoomName returns roomname of roomdata
func (e *ChatRoomOperator) GetRoomName() string {
	return e.roomData.RoomName
}

// GetIdentifier returns Identifier of roomdata
func (e *ChatRoomOperator) GetIdentifier() string {
	return e.roomData.Identifier
}

// ComparePassword compares input password to roomdata password
func (e *ChatRoomOperator) ComparePassword(password string) error {
	return PasswordOperator{}.ComparePassword(e.roomData.HashedPassword, password, e.roomData.Salt)
}

// GobEncode serialized roomdata to bytes
func (e *ChatRoomOperator) GobEncode() ([]byte, error) {
	w := new(bytes.Buffer)
	encoder := gob.NewEncoder(w)

	if err := encoder.Encode(e.roomData.Identifier); err != nil {
		return nil, err
	}
	if err := encoder.Encode(e.roomData.RoomName); err != nil {
		return nil, err
	}
	if err := encoder.Encode(e.roomData.HashedPassword); err != nil {
		return nil, err
	}
	if err := encoder.Encode(e.roomData.Salt); err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

// GobDecode decode bytes to roomdata
func (e *ChatRoomOperator) GobDecode(buf []byte) error {
	r := bytes.NewBuffer(buf)
	decoder := gob.NewDecoder(r)

	if err := decoder.Decode(&e.roomData.Identifier); err != nil {
		return err
	}
	if err := decoder.Decode(&e.roomData.RoomName); err != nil {
		return err
	}
	if err := decoder.Decode(&e.roomData.HashedPassword); err != nil {
		return err
	}
	if err := decoder.Decode(&e.roomData.Salt); err != nil {
		return err
	}
	return nil
}
