package models

import "time"

type User struct {
	Id          uint      `gorm:"primaryKey" json:"id"`
	Username    string    `gorm:"type:varchar(155);unique" json:"username" validate:"required"`
	Email       string    `gorm:"type:varchar(155);unique" json:"email" validate:"required,email"`
	Password    string    `gorm:"type:varchar(155)" json:"password" validate:"required,min=6"`
	Age         int       `gorm:"type:int(50)" json:"age" validate:"required,min=8"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Photos      []Photo
	Comments    []Comment
	SocialMedia []SocialMedia
}
