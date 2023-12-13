package repository

import (
	"context"
	"github.com/KaguraGateway/bingogame2-backend/domain/model"
)

type RoomUserRepository interface {
	FindByRoomID(ctx context.Context, roomID string) ([]*model.RoomUser, error)
	FindByRoomIDAndUserID(ctx context.Context, roomID string, userID string) (*model.RoomUser, error)
	Save(ctx context.Context, roomUser *model.RoomUser) error
}
