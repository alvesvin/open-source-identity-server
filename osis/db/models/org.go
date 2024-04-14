package models

import "gorm.io/gorm"

type Org struct {
	gorm.Model
	Name string
}
