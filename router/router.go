package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ordomigato/con-game/controller"
	"github.com/ordomigato/con-game/service"
)

var (
	roomService    service.RoomService       = service.New()
	RoomController controller.RoomController = controller.New(roomService)
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	api.GET("/join", func(c *gin.Context) {
		c.JSON(200, RoomController.Join(c))
	})
	api.POST("/create", func(c *gin.Context) {
		c.JSON(200, RoomController.Create(c))
	})

	return router

}
