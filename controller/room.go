package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ordomigato/con-game/entity"
	"github.com/ordomigato/con-game/service"
)

type RoomController interface {
	Create(ctx *gin.Context) entity.Room
	Join(ctx *gin.Context) entity.Room
}

type controller struct {
	service service.RoomService
}

func New(service service.RoomService) RoomController {
	return &controller{
		service: service,
	}
}

func (c *controller) Create(ctx *gin.Context) entity.Room {
	var room entity.Room
	ctx.BindJSON(&room)
	return c.service.Create(room)
}

func (c *controller) Join(ctx *gin.Context) entity.Room {
	var room entity.Room
	ctx.BindJSON(&room)
	return c.service.Join(room)
}
