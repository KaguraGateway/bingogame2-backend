package bundb

import (
	"context"
	"github.com/KaguraGateway/bingogame2-backend/domain/model"
	"github.com/KaguraGateway/bingogame2-backend/domain/repository"
	"github.com/KaguraGateway/bingogame2-backend/infra/bundb/dao"
	"github.com/samber/do"
	"github.com/samber/lo"
	"github.com/uptrace/bun"
)

type roomBingoNumberRepositoryDb struct {
	db *bun.DB
}

func NewRoomBingoNumberRepository(i *do.Injector) (repository.RoomBingoNumberRepository, error) {
	return roomBingoNumberRepositoryDb{
		db: do.MustInvoke[*bun.DB](i),
	}, nil
}

func (r roomBingoNumberRepositoryDb) FindByRoomID(ctx context.Context, roomID string) ([]*model.RoomBingoNumber, error) {
	roomBingoNumbers := make([]*dao.RoomBingoNumber, 0)
	if err := r.db.NewSelect().Model(&roomBingoNumbers).Where("room_id = ?", roomID).Scan(ctx); err != nil {
		return nil, err
	}
	return lo.Map(roomBingoNumbers, func(item *dao.RoomBingoNumber, index int) *model.RoomBingoNumber {
		return model.RebuildRoomBingoNumber(item.ID, roomID)
	}), nil
}

func (r roomBingoNumberRepositoryDb) Save(ctx context.Context, roomBingoNumber *model.RoomBingoNumber) error {
	daoRoomBingoNumber := &dao.RoomBingoNumber{
		ID:     roomBingoNumber.ID(),
		RoomID: roomBingoNumber.RoomID(),
	}
	if _, err := r.db.NewInsert().Model(daoRoomBingoNumber).Exec(ctx); err != nil {
		return err
	}
	return nil
}
