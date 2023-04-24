package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/ordomigato/con-game/controller"
	"github.com/ordomigato/con-game/middleware"
	"github.com/ordomigato/con-game/service"
)

var (
	gameService    service.GameService       = service.New()
	GameController controller.GameController = controller.New(gameService)
)

func SetupRouter(socketio *socketio.Server) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{"content-type"},
	}))
	router.Use(middleware.ErrorHandler)
	router.Use(middleware.SetHeaders)

	router.GET("/socket.io/*any", gin.WrapH(socketio))
	router.POST("/socket.io/*any", gin.WrapH(socketio))

	api := router.Group("/api")
	api.POST("/join", func(c *gin.Context) {
		GameController.Join(c)
	})
	api.POST("/create", func(c *gin.Context) {
		GameController.Create(c)
	})

	return router

}
