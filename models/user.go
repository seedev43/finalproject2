package models

import (
	"fp2/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	Id          uint      `gorm:"primaryKey" json:"id"`
	Username    string    `gorm:"type:varchar(155);unique" json:"username" valid:"required~Your username is required"`
	Email       string    `gorm:"type:varchar(155);unique" json:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password    string    `gorm:"type:varchar(155)" json:"password" valid:"required~Your password is required,minstringlength(6)~Minimum password length is 6 characters"`
	Age         int       `gorm:"type:int(50)" json:"age" valid:"required~Your age is required,range(8|99)~Minimum age 8 years"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Photos      []Photo
	Comments    []Comment
	SocialMedia []SocialMedia
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}
	u.Password = helpers.HashPass(u.Password)

	err = nil
	return
}
