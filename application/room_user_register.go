package application

import (
	"context"
	"github.com/KaguraGateway/bingogame2-backend/domain/model"
	"github.com/KaguraGateway/bingogame2-backend/domain/repository"
	"github.com/cockroachdb/errors"
	"github.com/samber/do"
)

type RoomUserRegisterInput struct {
	UserName string
	RoomId   string
}

type RoomUserRegisterOutput struct {
	BingoCard [5][5]model.BingoRow
	UserId    string
}

type RoomUserRegister interface {
	Execute(ctx context.Context, input RoomUserRegisterInput) (*RoomUserRegisterOutput, error)
}

type roomUserRegisterUseCase struct {
	roomRepository     repository.RoomRepository
	roomUserRepository repository.RoomUserRepository
}

func NewRoomUserRegisterUseCase(i *do.Injector) (RoomUserRegister, error) {
	return roomUserRegisterUseCase{
		roomRepository:     do.MustInvoke[repository.RoomRepository](i),
		roomUserRepository: do.MustInvoke[repository.RoomUserRepository](i),
	}, nil
}

func (r roomUserRegisterUseCase) Execute(ctx context.Context, input RoomUserRegisterInput) (*RoomUserRegisterOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, CtxTimeoutDur)
	defer cancel()

	room, err := r.roomRepository.FindByID(ctx, input.RoomId)
	if err != nil {
		return nil, errors.Join(err, ErrRoomNotFound)
	}

	roomUser := model.NewRoomUser(room.ID(), input.UserName)
	if err := r.roomUserRepository.Save(ctx, roomUser); err != nil {
		return nil, errors.Join(err, ErrFailedCreateRoomUser)
	}

	return &RoomUserRegisterOutput{
		BingoCard: model.GenerateCard(roomUser.BingoSeed()),
		UserId:    roomUser.ID(),
	}, nil
}
