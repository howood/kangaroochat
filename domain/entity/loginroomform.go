package entity

// LoginRoomForm entity
type LoginRoomForm struct {
	UserName string `validate:"required,min=8"`
	Password string `validate:"required"`
}
