package repository

import "github.com/KaguraGateway/bingogame2-backend/domain/model"

type RoomUserRepository interface {
	FindByRoomIDAndUserID(roomID string, userID string) (*model.RoomUser, error)
	Save(roomUser *model.RoomUser) error
}
