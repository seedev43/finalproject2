package dto

import "time"

type PhotoResponse struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserId   uint   `json:"user_id"`
}

type PhotosResponse struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"type:varchar(155)" json:"title" valid:"required~Your title is required"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url" valid:"required~Your photo_url is required"`
	UserId    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      UserResponse
}

type PhotoCreateResponse struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type PhotoUpdateResponse struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (PhotoResponse) TableName() string {
	return "photos"
}

func (PhotosResponse) TableName() string {
	return "photos"
}
