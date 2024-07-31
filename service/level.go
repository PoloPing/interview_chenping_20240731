package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/poloping/interview_chenping_20240731/data/request"
	"github.com/poloping/interview_chenping_20240731/data/response"
	"github.com/poloping/interview_chenping_20240731/helper"
	"github.com/poloping/interview_chenping_20240731/model"
	"github.com/poloping/interview_chenping_20240731/repository"
)

type LevelService interface {
	Create(player request.CreateLevelRequest) error
	FindAll() []response.LevelResponse
}

type LevelServiceImpl struct {
	LevelRepository repository.LevelRepository
	Validate        *validator.Validate
}

func NewLevelServiceImpl(LevelRepository repository.LevelRepository, validate *validator.Validate) LevelService {
	return &LevelServiceImpl{
		LevelRepository: LevelRepository,
		Validate:        validate,
	}
}

// Create implements LevelService
func (l *LevelServiceImpl) Create(level request.CreateLevelRequest) error {
	err := l.Validate.Struct(level)
	if err != nil {
		return err
	}
	helper.ErrorPanic(err)
	levelModel := model.Level{
		Name: level.Name,
	}
	l.LevelRepository.Save(levelModel)

	return nil
}

// FindAll implements LevelService
func (l *LevelServiceImpl) FindAll() []response.LevelResponse {
	result := l.LevelRepository.FindAll()

	var levels []response.LevelResponse
	for _, value := range result {
		level := response.LevelResponse{
			ID:   value.ID,
			Name: value.Name,
		}

		levels = append(levels, level)
	}

	return levels
}
