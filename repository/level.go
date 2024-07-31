package repository

import (
	"github.com/poloping/interview_chenping_20240731/helper"
	"github.com/poloping/interview_chenping_20240731/model"
	"gorm.io/gorm"
)

type LevelRepository interface {
	Save(level model.Level)
	FindAll() []model.Level
}

type LevelRepositoryImpl struct {
	Db *gorm.DB
}

func NewLevelRepositoryImpl(Db *gorm.DB) LevelRepository {
	return &LevelRepositoryImpl{Db: Db}
}

func (l LevelRepositoryImpl) Save(level model.Level) {
	result := l.Db.Create(&level)
	helper.ErrorPanic(result.Error)
}

func (l LevelRepositoryImpl) FindAll() []model.Level {
	var levels []model.Level
	results := l.Db.Find(&levels)
	helper.ErrorPanic(results.Error)
	return levels
}
