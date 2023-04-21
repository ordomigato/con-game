package service

import (
	"github.com/ordomigato/con-game/entity"
	"github.com/ordomigato/con-game/utils"
)

type GameService interface {
	Create(entity.GameConfig) (*entity.Game, error)
	// Join(entity.User) entity.Game
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
func (service *gameService) Create(gameConfig entity.GameConfig) (*entity.Game, error) {
	useCodes := utils.GetKeysFromMap(service.gameRooms)
	code := generateUniqueGameCode(useCodes)
	gameConfig.RoomCode = code
	var gameRoom entity.Game = entity.Game{
		GameConfig: gameConfig,
	}
	service.gameRooms[code] = gameRoom
	return &gameRoom, nil
}

// func (service *gameService) Join(roomCode string, user entity.User) entity.Game {
// find room
// add user to room and setup socket connection
// return game room
// }
