package entity

// ChatRoom entity
type ChatRoom struct {
	Identifier     string
	RoomName       string
	HashedPassword string
	Salt           string
}
