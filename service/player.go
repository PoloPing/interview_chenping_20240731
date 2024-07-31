package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/poloping/interview_chenping_20240731/data/request"
	"github.com/poloping/interview_chenping_20240731/data/response"
	"github.com/poloping/interview_chenping_20240731/helper"
	"github.com/poloping/interview_chenping_20240731/model"
	"github.com/poloping/interview_chenping_20240731/repository"
)

type PlayerService interface {
	Create(player request.CreatePlayerRequest) error
	Update(player request.UpdatePlayerRequest) error
	Delete(playerID int)
	FindById(playerID int) response.PlayerResponse
	FindAll() []response.PlayerResponse
}

type PlayerServiceImpl struct {
	PlayerRepository repository.PlayerRepository
	Validate         *validator.Validate
}

func NewPlayerServiceImpl(playerRepository repository.PlayerRepository, validate *validator.Validate) PlayerService {
	return &PlayerServiceImpl{
		PlayerRepository: playerRepository,
		Validate:         validate,
	}
}

// Create implements PlayerService
func (p *PlayerServiceImpl) Create(player request.CreatePlayerRequest) error {
	err := p.Validate.Struct(player)
	if err != nil {
		return err
	}
	helper.ErrorPanic(err)
	playerModel := model.Player{
		Name:     player.Name,
		PlayerID: player.PlayerID,
		LevelID:  player.LevelID,
	}
	p.PlayerRepository.Save(playerModel)

	return nil
}

// Delete implements PlayerService
func (p *PlayerServiceImpl) Delete(playerID int) {
	p.PlayerRepository.Delete(playerID)
}

// FindAll implements PlayerService
func (p *PlayerServiceImpl) FindAll() []response.PlayerResponse {
	result := p.PlayerRepository.FindAll()

	var players []response.PlayerResponse
	for _, value := range result {
		level := response.LevelResponse{
			ID:   value.Level.ID,
			Name: value.Level.Name,
		}
		player := response.PlayerResponse{
			ID:            value.ID,
			Name:          value.Name,
			PlayerID:      value.PlayerID,
			LevelResponse: level,
		}

		players = append(players, player)
	}

	return players
}

// FindById implements PlayerService
func (p *PlayerServiceImpl) FindById(playerID int) response.PlayerResponse {
	playerData, err := p.PlayerRepository.FindById(playerID)
	helper.ErrorPanic(err)

	level := response.LevelResponse{
		ID:   playerData.Level.ID,
		Name: playerData.Level.Name,
	}

	playerResponse := response.PlayerResponse{
		ID:            playerData.ID,
		Name:          playerData.Name,
		LevelResponse: level,
	}
	return playerResponse
}

// Update implements PlayerService
func (p *PlayerServiceImpl) Update(player request.UpdatePlayerRequest) error {
	playerData, err := p.PlayerRepository.FindById(player.ID)
	if err != nil {
		return err
	}
	playerData.Name = player.Name
	playerData.PlayerID = player.PlayerID
	playerData.LevelID = player.LevelID
	p.PlayerRepository.Update(playerData)
	return nil
}
