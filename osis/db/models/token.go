package models

import (
	"database/sql"
	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	SessionId        uint
	Session          Session `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AccessTokenHash  string  `gorm:"uniqueIndex"`
	RefreshTokenHash string  `gorm:"uniqueIndex"`
	ExpiresAt        sql.NullTime
}
