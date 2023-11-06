package models

type SocialMedia struct {
	Id             uint   `gorm:"primaryKey" json:"id"`
	Name           string `gorm:"type:varchar(155)" json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserId         uint   `json:"user_id"`
}
