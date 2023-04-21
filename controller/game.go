package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ordomigato/con-game/entity"
	"github.com/ordomigato/con-game/service"
)

type GameController interface {
	Create(ctx *gin.Context) (*entity.Game, error)
	// Join(ctx *gin.Context) entity.Game
}

type controller struct {
	service service.GameService
}

func New(service service.GameService) GameController {
	return &controller{
		service: service,
	}
}

func (c *controller) Create(ctx *gin.Context) (*entity.Game, error) {
	var gameConfig entity.GameConfig
	err := ctx.BindJSON(&gameConfig)
	if err != nil {
		return nil, err
	}
	game, err := c.service.Create(gameConfig)
	return game, err
}

// func (c *controller) Join(ctx *gin.Context) entity.Game {
// 	var gameConfig entity.GameConfig
// 	ctx.BindJSON(&gameConfig)
// 	return c.service.Join(gameConfig)
// }
