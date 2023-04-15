package service

import "github.com/ordomigato/con-game/entity"

type RoomService interface {
	Create(entity.Room) entity.Room
	Join(entity.Room) entity.Room
}

type roomService struct {
	rooms []entity.Room
}

func New() RoomService {
	return &roomService{}
}

func (service *roomService) Create(room entity.Room) entity.Room {
	service.rooms = append(service.rooms, room)
	return room
}

func (service *roomService) Join(room entity.Room) entity.Room {
	return room
}
