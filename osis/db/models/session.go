package models

import (
	"database/sql"
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	UserId       uint
	User         User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AccountId    uint
	Account      Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Ip           string
	DeviceName   string
	TrustedUntil sql.NullTime
}
