package repository

import (
	"context"
	"github.com/KaguraGateway/bingogame2-backend/domain/model"
)

type RoomExchangedPrizeRepository interface {
	FindByRoomID(ctx context.Context, roomID string) ([]*model.RoomExchangedPrize, error)
	Save(ctx context.Context, roomExchangedPrize *model.RoomExchangedPrize) error
}
