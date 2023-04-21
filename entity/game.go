package entity

type Game struct {
	GameConfig GameConfig `json:"gameConfig"`
	Users      []User     `json:"users"`
}

type GameConfig struct {
	RoomCode string `json:"roomCode"`
}
