package models

import (
	"fp2/dto"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	Id             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"type:varchar(155)" json:"name" valid:"required~Your name is required"`
	SocialMediaUrl string    `json:"social_media_url" valid:"required~Your social media url is required"`
	UserId         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	User           dto.UserSocialMediaResponse
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}

func (s *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
