package model

import "gorm.io/gorm"

type Level struct {
	gorm.Model
	Name string
}
