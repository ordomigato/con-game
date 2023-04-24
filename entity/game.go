package entity

type Game struct {
	GameConfig GameConfig `json:"gameConfig"`
	Players    []Players  `json:"players"`
}

type GameConfig struct {
	RoomCode   string `json:"roomCode"`
	MaxPlayers int    `json:"maxPlayers" binding:"required"`
}
