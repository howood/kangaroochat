package actor

import (
	"bytes"
	"context"
	"encoding/gob"

	"github.com/howood/kangaroochat/domain/entity"
	"github.com/howood/kangaroochat/domain/repository"
	"github.com/howood/kangaroochat/infrastructure/uuid"
)

// chatRoom struct
type ChatRoomOperator struct {
	repository.ChatRoomRepository
}

// NewChatRoomOperator creates a new ChatRoomRepository
func NewChatRoomOperator(ctx context.Context) *ChatRoomOperator {
	return &ChatRoomOperator{&chatRoom{ctx: ctx}}
}

// chatRoom struct
type chatRoom struct {
	roomData entity.ChatRoom
	ctx      context.Context
}

// Set sets roomname and password to  roomdata
func (e *chatRoom) Set(roomname, password string) error {
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
func (e *chatRoom) GetRoomName() string {
	return e.roomData.RoomName
}

// GetIdentifier returns Identifier of roomdata
func (e *chatRoom) GetIdentifier() string {
	return e.roomData.Identifier
}

// ComparePassword compares input password to roomdata password
func (e *chatRoom) ComparePassword(password string) error {
	return PasswordOperator{}.ComparePassword(e.roomData.HashedPassword, password, e.roomData.Salt)
}

// GobEncode serialized roomdata to bytes
func (e *chatRoom) GobEncode() ([]byte, error) {
	w := new(bytes.Buffer)
	encoder := gob.NewEncoder(w)

	var err error
	if err == nil {
		err = encoder.Encode(e.roomData.Identifier)
	}
	if err == nil {
		err = encoder.Encode(e.roomData.RoomName)
	}
	if err == nil {
		err = encoder.Encode(e.roomData.HashedPassword)
	}
	if err == nil {
		err = encoder.Encode(e.roomData.Salt)
	}
	if err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

// GobDecode decode bytes to roomdata
func (e *chatRoom) GobDecode(buf []byte) error {
	r := bytes.NewBuffer(buf)
	decoder := gob.NewDecoder(r)

	var err error
	if err == nil {
		err = decoder.Decode(&e.roomData.Identifier)
	}
	if err == nil {
		err = decoder.Decode(&e.roomData.RoomName)
	}
	if err == nil {
		err = decoder.Decode(&e.roomData.HashedPassword)
	}
	if err == nil {
		err = decoder.Decode(&e.roomData.Salt)
	}
	return err
}
