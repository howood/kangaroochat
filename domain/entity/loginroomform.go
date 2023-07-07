package entity

// LoginRoomForm entity
type LoginRoomForm struct {
	UserName string `form:"username" validate:"required,min=8"`
	Password string `form:"password" validate:"required"`
}
