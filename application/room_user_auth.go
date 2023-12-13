package application

import (
	"context"
	"github.com/KaguraGateway/bingogame2-backend/domain/model"
	"github.com/KaguraGateway/bingogame2-backend/domain/repository"
	"github.com/cockroachdb/errors"
	"github.com/samber/do"
	"github.com/samber/lo"
)

type RoomUserAuthInput struct {
	UserId string
	RoomId string
}

type RoomUserAuthOutput struct {
	bingoCard [5][5]model.BingoRow
	userId    string
}

type RoomUserAuth interface {
	Execute(ctx context.Context, input RoomUserAuthInput) (*RoomUserAuthOutput, error)
}

type roomUserAuthUseCase struct {
	roomRepository            repository.RoomRepository
	roomBingoNumberRepository repository.RoomBingoNumberRepository
	roomUserRepository        repository.RoomUserRepository
}

func NewRoomUserAuthUseCase(i *do.Injector) (RoomUserAuth, error) {
	return roomUserAuthUseCase{
		roomRepository:            do.MustInvoke[repository.RoomRepository](i),
		roomBingoNumberRepository: do.MustInvoke[repository.RoomBingoNumberRepository](i),
		roomUserRepository:        do.MustInvoke[repository.RoomUserRepository](i),
	}, nil
}

func (r roomUserAuthUseCase) Execute(ctx context.Context, input RoomUserAuthInput) (*RoomUserAuthOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, CtxTimeoutDur)
	defer cancel()

	room, err := r.roomRepository.FindByID(ctx, input.RoomId)
	if err != nil {
		return nil, errors.Join(err, ErrRoomNotFound)
	}

	roomBingoNumbers, err := r.roomBingoNumberRepository.FindByRoomID(ctx, room.ID())
	if err != nil {
		return nil, errors.Join(err, ErrRoomNotFound)
	}

	roomUser, err := r.roomUserRepository.FindByRoomIDAndUserID(ctx, room.ID(), input.UserId)
	if err != nil {
		return nil, errors.Join(err, ErrRoomUserNotFound)
	}

	return &RoomUserAuthOutput{
		bingoCard: roomUser.BingoCard(lo.Map(roomBingoNumbers, func(roomBingoNumber *model.RoomBingoNumber, _ int) uint {
			return roomBingoNumber.ID()
		})).Card(),
		userId: roomUser.ID(),
	}, nil
}
