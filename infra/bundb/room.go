package bundb

import (
	"context"
	"github.com/KaguraGateway/bingogame2-backend/domain/model"
	"github.com/KaguraGateway/bingogame2-backend/domain/repository"
	"github.com/KaguraGateway/bingogame2-backend/infra/bundb/dao"
	"github.com/samber/do"
	"github.com/uptrace/bun"
)

type roomRepositoryDb struct {
	db *bun.DB
}

func NewRoomRepository(i *do.Injector) (repository.RoomRepository, error) {
	return roomRepositoryDb{
		db: do.MustInvoke[*bun.DB](i),
	}, nil
}

func (r roomRepositoryDb) FindByID(ctx context.Context, id string) (*model.Room, error) {
	room := new(model.Room)
	if err := r.db.NewSelect().Model(&room).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, err
	}
	return room, nil
}

func (r roomRepositoryDb) Save(ctx context.Context, room *model.Room) error {
	daoRoom := &dao.Room{
		ID:       room.ID(),
		PrizeNum: room.PrizeNum(),
	}
	if _, err := r.db.NewInsert().Model(daoRoom).Exec(ctx); err != nil {
		return err
	}
	return nil
}
