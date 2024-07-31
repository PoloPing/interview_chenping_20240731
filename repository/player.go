package repository

import (
	"errors"
	"github.com/poloping/interview_chenping_20240731/data/request"
	"github.com/poloping/interview_chenping_20240731/helper"
	"github.com/poloping/interview_chenping_20240731/model"
	"gorm.io/gorm"
)

type PlayerRepository interface {
	Save(player model.Player)
	Update(player model.Player)
	Delete(playerId int)
	FindById(playerId int) (player model.Player, err error)
	FindAll() []model.Player
}

type PlayerRepositoryImpl struct {
	Db *gorm.DB
}

func NewPlayerRepositoryImpl(Db *gorm.DB) PlayerRepository {
	return &PlayerRepositoryImpl{Db: Db}
}

func (p PlayerRepositoryImpl) Save(player model.Player) {
	result := p.Db.Create(&player)
	helper.ErrorPanic(result.Error)
}

func (p PlayerRepositoryImpl) Update(player model.Player) {
	var updatePlayer = request.UpdatePlayerRequest{
		ID:       int(player.ID),
		Name:     player.Name,
		PlayerID: player.PlayerID,
		LevelID:  player.LevelID,
	}
	result := p.Db.Model(&player).Updates(updatePlayer)
	helper.ErrorPanic(result.Error)
}

func (p PlayerRepositoryImpl) Delete(playerId int) {
	var player model.Player
	result := p.Db.Where("id = ?", playerId).Delete(&player)
	helper.ErrorPanic(result.Error)
}

func (p PlayerRepositoryImpl) FindById(playerId int) (model.Player, error) {
	var player model.Player
	result := p.Db.Preload("Level").Find(&player, playerId)
	if result != nil {
		return player, nil
	} else {
		return player, errors.New("player is not found")
	}
}

func (p PlayerRepositoryImpl) FindAll() []model.Player {
	var players []model.Player
	results := p.Db.Preload("Level").Find(&players)
	helper.ErrorPanic(results.Error)
	return players
}
