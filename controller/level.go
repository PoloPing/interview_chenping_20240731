package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/poloping/interview_chenping_20240731/data/request"
	"github.com/poloping/interview_chenping_20240731/data/response"
	"github.com/poloping/interview_chenping_20240731/service"
	"github.com/rs/zerolog/log"
	"net/http"
)

type LevelController struct {
	levelService service.LevelService
}

func NewLevelController(service service.LevelService) *LevelController {
	return &LevelController{
		levelService: service,
	}
}

// Create CreateLevel		godoc
// @Summary			Create level
// @Description		Save level data in Db.
// @Param			level body request.CreateLevelRequest true "Create level"
// @Produce			application/json
// @Tags			level
// @Success			200 {object} response.Response{}
// @Router			/levels [post]
func (controller *LevelController) Create(ctx *gin.Context) {
	log.Info().Msg("create level")
	createLevelRequest := request.CreateLevelRequest{}
	err := ctx.ShouldBindJSON(&createLevelRequest)

	if err != nil {
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "status": err, "data": gin.H{}})
		return
	}

	err = controller.levelService.Create(createLevelRequest)
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

// FindAll FindAllLevel		godoc
// @Summary			Get All level.
// @Description		Return list of level.
// @Tags			level
// @Success			200 {object} response.Response{}
// @Router			/levels [get]
func (controller *LevelController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll levels")
	levelResponse := controller.levelService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   levelResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
