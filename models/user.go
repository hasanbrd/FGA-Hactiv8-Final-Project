package models

import (
	"final-project/helpers"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Username     string `gorm:"not null;uniqueIndex"`
	Email        string `gorm:"not null;uniqueIndex" json:"email"`
	Password     string `gorm:"not null"`
	Age          uint   `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Comments     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SocialMedias []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.Password = helpers.HashPass(u.Password)

	err = nil
	return
}
