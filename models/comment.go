package models

import (
	"fp2/dto"
	"time"
)

type Comment struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	UserId    uint      `json:"user_id"`
	PhotoId   uint      `json:"photo_id"`
	Message   string    `json:"message" valid:"required~Your message is required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      dto.UserResponse
	Photo     dto.PhotoResponse
}
