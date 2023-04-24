package service

import (
	"fmt"
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

func (service *gameService) Join(ctx *gin.Context, roomCode string, playerName string) {
	gameRoom, ok := service.gameRooms[roomCode]
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"errors": "no room found",
		})
		return
	}
	for _, p := range gameRoom.Players {
		if p.Name == playerName {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"errors": "username already taken",
			})
			return
		}
	}
	if len(gameRoom.Players) == gameRoom.GameConfig.MaxPlayers {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"errors": "room is full",
		})
		return
	}
	gameRoom.Players = append(
		service.gameRooms[roomCode].Players,
		entity.Players{
			Name: playerName,
		},
	)
	service.gameRooms[roomCode] = gameRoom
	fmt.Println(gameRoom.Players)
	ctx.JSON(http.StatusOK, gameRoom)
}
