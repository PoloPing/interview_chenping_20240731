package model

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	Name     string
	PlayerID string
	LevelID  int   // Foreign Key
	Level    Level // referenced Model
}
