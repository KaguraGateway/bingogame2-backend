package application

import (
	"context"
	"github.com/KaguraGateway/bingogame2-backend/domain/model"
	"github.com/KaguraGateway/bingogame2-backend/domain/repository"
	"github.com/samber/do"
)

type CreateRoomInput struct {
	PrizeNum uint
}

type CreateRoomOutput struct {
	RoomId string
}

type CreateRoom interface {
	Execute(ctx context.Context, input CreateRoomInput) (*CreateRoomOutput, error)
}

type createRoomUseCase struct {
	roomRepository repository.RoomRepository
}

func NewCreateRoomUseCase(i *do.Injector) (CreateRoom, error) {
	return createRoomUseCase{
		roomRepository: do.MustInvoke[repository.RoomRepository](i),
	}, nil
}

func (r createRoomUseCase) Execute(ctx context.Context, input CreateRoomInput) (*CreateRoomOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, CtxTimeoutDur)
	defer cancel()

	room := model.NewRoom(input.PrizeNum)
	if err := r.roomRepository.Save(ctx, room); err != nil {
		return nil, err
	}

	return &CreateRoomOutput{
		RoomId: room.ID(),
	}, nil
}
