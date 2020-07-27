package entity

// CreateRoomForm entity
type CreateRoomForm struct {
	RoomName string `validate:"required"`
	Password string `validate:"required,min=8"`
}
