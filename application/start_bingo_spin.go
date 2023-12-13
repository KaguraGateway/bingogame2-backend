package application

import (
	"context"
	"github.com/KaguraGateway/bingogame2-backend/domain/model"
	"github.com/KaguraGateway/bingogame2-backend/domain/repository"
	"github.com/samber/do"
	"github.com/samber/lo"
	"math/rand"
)

type StartBingoSpinInput struct {
	RoomId string
}

type StartBingoSpinOutput struct {
	BingoNumber uint
	BingoUsers  []*model.RoomUser
	BingoCards  map[string]*model.BingoCard
}

type StartBingoSpin interface {
	Execute(ctx context.Context, input StartBingoSpinInput) (*StartBingoSpinOutput, error)
}

type startBingoSpinUseCase struct {
	roomRepository            repository.RoomRepository
	roomUserRepository        repository.RoomUserRepository
	roomBingoNumberRepository repository.RoomBingoNumberRepository
}

func NewStartBingoSpinUseCase(i *do.Injector) (StartBingoSpin, error) {
	return startBingoSpinUseCase{
		roomRepository:            do.MustInvoke[repository.RoomRepository](i),
		roomUserRepository:        do.MustInvoke[repository.RoomUserRepository](i),
		roomBingoNumberRepository: do.MustInvoke[repository.RoomBingoNumberRepository](i),
	}, nil
}

func (r startBingoSpinUseCase) Execute(ctx context.Context, input StartBingoSpinInput) (*StartBingoSpinOutput, error) {
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
	bingoNumbers := lo.Map(roomBingoNumbers, func(roomBingoNumber *model.RoomBingoNumber, _ int) uint {
		return roomBingoNumber.ID()
	})

	// ドメインモデル貧血症なので解消したい
	allBingoNumbers := lo.RangeFrom(1, 75)
	bingoNumberSeeds := lo.Filter(allBingoNumbers, func(number int, _ int) bool {
		return !lo.Contains(bingoNumbers, uint(number))
	})
	bingoNumber := uint(bingoNumberSeeds[rand.Intn(len(bingoNumberSeeds))])
	bingoNumbers = append(bingoNumbers, bingoNumber)

	roomBingoNumber := model.NewRoomBingoNumber(bingoNumber, room.ID())
	if err := r.roomBingoNumberRepository.Save(ctx, roomBingoNumber); err != nil {
		return nil, err
	}

	// ユーザーがビンゴしたかどうか
	roomUsers, err := r.roomUserRepository.FindByRoomID(ctx, room.ID())
	if err != nil {
		return nil, err
	}

	var bingoUsers []*model.RoomUser
	var bingoCards map[string]*model.BingoCard
	for _, roomUser := range roomUsers {
		if roomUser.IsExchanged {
			continue
		}
		bingoCard := roomUser.BingoCard(bingoNumbers)
		if bingoCard.CheckBingo() {
			bingoUsers = append(bingoUsers, roomUser)
		}
		bingoCards[roomUser.ID()] = bingoCard
	}

	return &StartBingoSpinOutput{
		BingoNumber: bingoNumber,
		BingoUsers:  bingoUsers,
		BingoCards:  bingoCards,
	}, nil
}
