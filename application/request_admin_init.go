package application

import (
	"context"
	"github.com/KaguraGateway/bingogame2-backend/domain/model"
	"github.com/KaguraGateway/bingogame2-backend/domain/repository"
	"github.com/samber/do"
)

type RequestAdminInitInput struct {
	RoomId string
}

type RequestAdminInitOutput struct {
	bingoNumbers    []*model.RoomBingoNumber
	exchangedPrizes []*model.RoomExchangedPrize
	prizeNum        uint
}

type RequestAdminInit interface {
	Execute(ctx context.Context, input RequestAdminInitInput) (*RequestAdminInitOutput, error)
}

type requestAdminInitUseCase struct {
	roomRepository               repository.RoomRepository
	roomBingoNumberRepository    repository.RoomBingoNumberRepository
	roomExchangedPrizeRepository repository.RoomExchangedPrizeRepository
}

func NewRequestAdminInitUseCase(i *do.Injector) (RequestAdminInit, error) {
	return requestAdminInitUseCase{
		roomRepository:               do.MustInvoke[repository.RoomRepository](i),
		roomBingoNumberRepository:    do.MustInvoke[repository.RoomBingoNumberRepository](i),
		roomExchangedPrizeRepository: do.MustInvoke[repository.RoomExchangedPrizeRepository](i),
	}, nil
}

func (r requestAdminInitUseCase) Execute(ctx context.Context, input RequestAdminInitInput) (*RequestAdminInitOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, CtxTimeoutDur)
	defer cancel()

	room, err := r.roomRepository.FindByID(ctx, input.RoomId)
	if err != nil {
		return nil, err
	}

	roomBingoNumbers, err := r.roomBingoNumberRepository.FindByRoomID(ctx, room.ID())
	if err != nil {
		return nil, err
	}

	roomExchangedPrizes, err := r.roomExchangedPrizeRepository.FindByRoomID(ctx, room.ID())
	if err != nil {
		return nil, err
	}

	return &RequestAdminInitOutput{
		bingoNumbers:    roomBingoNumbers,
		exchangedPrizes: roomExchangedPrizes,
		prizeNum:        room.PrizeNum(),
	}, nil
}
