package dto

import "time"

type UserResponse struct {
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserSocialMediaResponse struct {
	Id              uint   `json:"id"`
	Username        string `json:"username"`
	ProfileImageUrl string `json:"profile_image_url"`
}

type UserRegsiterResponse struct {
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Age      int    `json:"age"`
}

type UserUpdateResponse struct {
	Id        uint      `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Age       int       `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (UserResponse) TableName() string {
	return "users"
}

func (UserSocialMediaResponse) TableName() string {
	return "users"
}
