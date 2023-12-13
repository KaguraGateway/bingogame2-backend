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

type roomUserRepositoryDb struct {
	db *bun.DB
}

func NewRoomUserRepository(i *do.Injector) (repository.RoomUserRepository, error) {
	return roomUserRepositoryDb{
		db: do.MustInvoke[*bun.DB](i),
	}, nil
}

func (r roomUserRepositoryDb) FindByRoomID(ctx context.Context, roomID string) ([]*model.RoomUser, error) {
	roomUsers := make([]*dao.RoomUser, 0)
	if err := r.db.NewSelect().Model(&roomUsers).Where("room_id = ?", roomID).Scan(ctx); err != nil {
		return nil, err
	}
	return lo.Map(roomUsers, func(roomUser *dao.RoomUser, index int) *model.RoomUser {
		return model.RebuildRoomUser(roomUser.ID, roomID, roomUser.Name, roomUser.BingoSeed, roomUser.IsExchanged)
	}), nil
}

func (r roomUserRepositoryDb) FindByRoomIDAndUserID(ctx context.Context, roomID string, userID string) (*model.RoomUser, error) {
	roomUser := new(dao.RoomUser)
	if err := r.db.NewSelect().Model(roomUser).Where("room_id = ?", roomID).Where("user_id = ?", userID).Scan(ctx); err != nil {
		return nil, err
	}
	return model.RebuildRoomUser(roomUser.ID, roomID, roomUser.Name, roomUser.BingoSeed, roomUser.IsExchanged), nil
}

func (r roomUserRepositoryDb) Save(ctx context.Context, roomUser *model.RoomUser) error {
	daoRoomUser := &dao.RoomUser{
		ID:     roomUser.ID(),
		RoomID: roomUser.RoomID(),
	}
	if _, err := r.db.NewInsert().Model(daoRoomUser).Exec(ctx); err != nil {
		return err
	}
	return nil
}
