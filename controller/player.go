package controller

import (
	"github.com/poloping/interview_chenping_20240731/data/request"
	"github.com/poloping/interview_chenping_20240731/data/response"
	"github.com/poloping/interview_chenping_20240731/helper"
	"github.com/poloping/interview_chenping_20240731/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type PlayerController struct {
	playerService service.PlayerService
}

func NewPlayerController(service service.PlayerService) *PlayerController {
	return &PlayerController{
		playerService: service,
	}
}

// Create CreatePlayer		godoc
// @Summary			Create player
// @Description		Save player data in Db.
// @Param			player body request.CreatePlayerRequest true "Create player"
// @Produce			application/json
// @Tags			player
// @Success			200 {object} response.Response{}
// @Router			/players [post]
func (controller *PlayerController) Create(ctx *gin.Context) {
	log.Info().Msg("create player")
	createPlayerRequest := request.CreatePlayerRequest{}
	err := ctx.ShouldBindJSON(&createPlayerRequest)

	if err != nil {
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "status": err, "data": gin.H{}})
		return
	}

	err = controller.playerService.Create(createPlayerRequest)
	if err != nil {
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "status": err.Error(), "data": gin.H{}})
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// Update UpdatePlayer		godoc
// @Summary			Update player
// @Description		Update player data.
// @Param			playerId path string true "update player by id"
// @Param			player body request.CreatePlayerRequest true  "Update player"
// @Produce			application/json
// @Tags			player
// @Success			200 {object} response.Response{}
// @Router			/players/{id} [put]
func (controller *PlayerController) Update(ctx *gin.Context) {
	log.Info().Msg("update player")

	ctx.Header("Content-Type", "application/json")

	updatePlayerRequest := request.UpdatePlayerRequest{}
	err := ctx.ShouldBindJSON(&updatePlayerRequest)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "status": err.Error(), "data": gin.H{}})
		return
	}

	playerID := ctx.Param("id")
	id, err := strconv.Atoi(playerID)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "status": err.Error(), "data": gin.H{}})
		return
	}

	updatePlayerRequest.ID = id

	err = controller.playerService.Update(updatePlayerRequest)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "status": err.Error(), "data": gin.H{}})
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   gin.H{},
	}

	ctx.JSON(http.StatusOK, webResponse)
}

// Delete DeletePlayer		godoc
// @Summary			Delete player
// @Description		Remove player data by id.
// @Produce			application/json
// @Tags			player
// @Success			200 {object} response.Response{}
// @Router			/player/{id} [delete]
func (controller *PlayerController) Delete(ctx *gin.Context) {
	log.Info().Msg("delete player")
	playerID := ctx.Param("id")
	id, err := strconv.Atoi(playerID)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "status": err.Error(), "data": gin.H{}})
		return
	}
	controller.playerService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindById FindByPlayerID		godoc
// @Summary				Get Single player by id.
// @Param				player path string true "update player by id"
// @Description			Return the player whose id value match id.
// @Produce				application/json
// @Tags				player
// @Success				200 {object} response.Response{}
// @Router				/players/{id} [get]
func (controller *PlayerController) FindById(ctx *gin.Context) {
	log.Info().Msg("findbyid player")
	playerID := ctx.Param("id")
	id, err := strconv.Atoi(playerID)
	helper.ErrorPanic(err)

	playerResponse := controller.playerService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   playerResponse,
	}
	ctx.Header("Content-Type", "application/json")
	if playerResponse.ID != 0 {
		ctx.JSON(http.StatusOK, webResponse)
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "status": "Not Found", "data": gin.H{}})
	}

}

// FindAll FindAllPlayer		godoc
// @Summary			Get All players.
// @Description		Return list of player.
// @Tags			player
// @Success			200 {object} response.Response{}
// @Router			/players [get]
func (controller *PlayerController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll player")
	playerResponse := controller.playerService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   playerResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
