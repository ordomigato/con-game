package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ordomigato/con-game/entity"
	"github.com/ordomigato/con-game/service"
)

type JoinRequestBody struct {
	PlayerName string `json:"name" binding:"required"`
	GameCode   string `json:"gameCode" binding:"required"`
}

type CreateRequestBody struct {
	GameConfig entity.GameConfig `json:"gameConfig" binding:"required"`
}

type GameController interface {
	Create(ctx *gin.Context)
	Join(ctx *gin.Context)
}

type controller struct {
	service service.GameService
}

func New(service service.GameService) GameController {
	return &controller{
		service: service,
	}
}

func (c *controller) Create(ctx *gin.Context) {
	reqBody := CreateRequestBody{}
	err := ctx.BindJSON(&reqBody)
	if err != nil {
		return
	}
	c.service.Create(ctx, reqBody.GameConfig)
}

func (c *controller) Join(ctx *gin.Context) {
	reqBody := JoinRequestBody{}
	err := ctx.BindJSON(&reqBody)
	if err != nil {
		return
	}
	c.service.Join(ctx, reqBody.GameCode, reqBody.PlayerName)
}
