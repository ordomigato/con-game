package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ordomigato/con-game/entity"
	"github.com/ordomigato/con-game/utils"
)

type GameService interface {
	Create(ctx *gin.Context, gameConfig entity.GameConfig)
	Join(ctx *gin.Context, roomCode string, username string)
}

type gameService struct {
	gameRooms map[string]entity.Game
}

func New() GameService {
	return &gameService{
		gameRooms: make(map[string]entity.Game),
	}
}

func generateUniqueGameCode(usedCodes []string) string {
	var code string
	for ok := true; ok; ok = utils.StringSearchSlice(usedCodes, code) {
		code = utils.GenerateCode(4)
	}
	return code
}

// Create game room
func (service *gameService) Create(ctx *gin.Context, gameConfig entity.GameConfig) {
	useCodes := utils.GetKeysFromMap(service.gameRooms)
	code := generateUniqueGameCode(useCodes)
	gameConfig.RoomCode = code
	var gameRoom entity.Game = entity.Game{
		GameConfig: gameConfig,
	}
	service.gameRooms[code] = gameRoom
	ctx.JSON(http.StatusOK, gameRoom)
}

func (service *gameService) Join(ctx *gin.Context, roomCode string, username string) {
	gameRoom, ok := service.gameRooms[roomCode]
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"errors": "no room found",
		})
		return
	}
	gameRoom.Users = append(
		gameRoom.Users,
		entity.User{
			Name: username,
		},
	)
	ctx.JSON(http.StatusOK, gameRoom)
}
