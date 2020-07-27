package repository

// ChatRoomRepository interface
type ChatRoomRepository interface {
	Set(roomname, password string) error
	GetRoomName() string
	GetIdentifier() string
	ComparePassword(password string) error
	GobEncode() ([]byte, error)
	GobDecode(buf []byte) error
}
