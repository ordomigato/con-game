package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ordomigato/con-game/controller"
	"github.com/ordomigato/con-game/middleware"
	"github.com/ordomigato/con-game/service"
)

var (
	gameService    service.GameService       = service.New()
	GameController controller.GameController = controller.New(gameService)
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.ErrorHandler)

	api := router.Group("/api")
	api.GET("/join", func(c *gin.Context) {
		GameController.Join(c)
	})
	api.POST("/create", func(c *gin.Context) {
		GameController.Create(c)
	})

	return router

}
