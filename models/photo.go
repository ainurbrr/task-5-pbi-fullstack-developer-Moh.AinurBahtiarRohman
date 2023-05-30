package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey" json:"id"`
	Title    string `json:"title" form:"title" valid:"required~Photo Title is Required"`
	Caption  string `json:"caption" form:"caption"`
	PhotoURL string `json:"photo_url" form:"photo_url" valid:"required~Photo URL is Required"`
	UserID   uint    `json:"user_id"`
	User     *User
}

func (u *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)

	if err != nil {
		return err
	}

	return nil
}

func (u *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)

	if err != nil {
		return err
	}

	return err
}
