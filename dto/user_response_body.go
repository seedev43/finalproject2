package dto

type UserResponse struct {
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func (UserResponse) TableName() string {
	return "users"
}
