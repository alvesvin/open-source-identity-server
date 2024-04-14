package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	UserId uint
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
