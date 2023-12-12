package repository

import (
	"context"
	"github.com/KaguraGateway/bingogame2-backend/domain/model"
)

type RoomRepository interface {
	FindByID(ctx context.Context, id string) (*model.Room, error)
	Save(ctx context.Context, room *model.Room) error
}
