package models

import "time"

type Photo struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"type:varchar(155)" json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
