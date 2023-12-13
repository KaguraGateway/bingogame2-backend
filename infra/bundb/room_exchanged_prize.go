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

type roomExchangedPrizeRepositoryDb struct {
	db *bun.DB
}

func NewRoomExchangedPrizeRepository(i *do.Injector) (repository.RoomExchangedPrizeRepository, error) {
	return roomExchangedPrizeRepositoryDb{
		db: do.MustInvoke[*bun.DB](i),
	}, nil
}

func (r roomExchangedPrizeRepositoryDb) FindByRoomID(ctx context.Context, roomID string) ([]*model.RoomExchangedPrize, error) {
	roomExchangedPrizes := make([]*dao.RoomExchangedPrize, 0)
	if err := r.db.NewSelect().Model(&roomExchangedPrizes).Where("room_id = ?", roomID).Scan(ctx); err != nil {
		return nil, err
	}
	return lo.Map(roomExchangedPrizes, func(item *dao.RoomExchangedPrize, index int) *model.RoomExchangedPrize {
		return model.RebuildRoomExchangedPrize(item.ID, roomID)
	}), nil
}

func (r roomExchangedPrizeRepositoryDb) Save(ctx context.Context, roomExchangedPrize *model.RoomExchangedPrize) error {
	daoRoomExchangedPrize := &dao.RoomExchangedPrize{
		ID:     roomExchangedPrize.ID(),
		RoomID: roomExchangedPrize.RoomID(),
	}
	if _, err := r.db.NewInsert().Model(daoRoomExchangedPrize).Exec(ctx); err != nil {
		return err
	}
	return nil
}
