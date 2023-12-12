package repository

import (
	"context"
	"github.com/KaguraGateway/bingogame2-backend/domain/model"
)

type RoomBingoNumberRepository interface {
	FindByRoomID(ctx context.Context, roomID string) ([]*model.RoomBingoNumber, error)
	Save(ctx context.Context, roomBingoNumber *model.RoomBingoNumber) error
}
