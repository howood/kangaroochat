package entity

// CreateRoomForm entity
type CreateRoomForm struct {
	RoomName string `form:"roomname" validate:"required"`
	Password string `form:"password" validate:"required,min=8"`
}
