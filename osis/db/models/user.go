package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string
	Email string
	Phone string
	OrgId *int
	Org   Org `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
