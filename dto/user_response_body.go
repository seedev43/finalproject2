package dto

type UserRegistered struct {
	Id       uint   `json:"id"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Age      int    `json:"age" validate:"required,min=8"`
}
