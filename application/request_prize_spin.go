package application

import (
	"context"
	"github.com/KaguraGateway/bingogame2-backend/domain/model"
	"github.com/KaguraGateway/bingogame2-backend/domain/repository"
	"github.com/samber/do"
	"github.com/samber/lo"
	"math/rand"
)

type RequestPrizeSpinInput struct {
	RoomId string
	UserId string
}

type RequestPrizeSpinOutput struct {
	PrizeNumber uint
	UserId      string
	UserName    string
	BingoUsers  []*model.RoomUser
}

type RequestPrizeSpin interface {
	Execute(ctx context.Context, input RequestPrizeSpinInput) (*RequestPrizeSpinOutput, error)
}

type requestPrizeSpinUseCase struct {
	roomRepo               repository.RoomRepository
	roomUserRepo           repository.RoomUserRepository
	roomBingoNumberRepo    repository.RoomBingoNumberRepository
	roomExchangedPrizeRepo repository.RoomExchangedPrizeRepository
}

func NewRequestPrizeSpinUseCase(i *do.Injector) (RequestPrizeSpin, error) {
	return requestPrizeSpinUseCase{
		roomRepo:               do.MustInvoke[repository.RoomRepository](i),
		roomUserRepo:           do.MustInvoke[repository.RoomUserRepository](i),
		roomBingoNumberRepo:    do.MustInvoke[repository.RoomBingoNumberRepository](i),
		roomExchangedPrizeRepo: do.MustInvoke[repository.RoomExchangedPrizeRepository](i),
	}, nil
}

func (r requestPrizeSpinUseCase) Execute(ctx context.Context, input RequestPrizeSpinInput) (*RequestPrizeSpinOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, CtxTimeoutDur)
	defer cancel()

	room, err := r.roomRepo.FindByID(ctx, input.RoomId)
	if err != nil {
		return nil, err
	}

	roomExchangedPrizes, err := r.roomExchangedPrizeRepo.FindByRoomID(ctx, room.ID())
	if err != nil {
		return nil, err
	}
	exchangedPrizes := lo.Map(roomExchangedPrizes, func(roomExchangedPrize *model.RoomExchangedPrize, _ int) uint {
		return roomExchangedPrize.ID()
	})

	roomUser, err := r.roomUserRepo.FindByRoomIDAndUserID(ctx, room.ID(), input.UserId)
	if err != nil {
		return nil, err
	} else if roomUser.IsExchanged {
		return nil, ErrAlreadyExchanged
	}

	// ドメインモデル貧血症なので解消したい
	allPrizes := lo.RangeFrom(1, int(room.PrizeNum()))
	stillPrizes := lo.Filter(allPrizes, func(prize int, _ int) bool {
		return lo.Contains(exchangedPrizes, uint(prize))
	})
	prizeNumber := uint(stillPrizes[rand.Intn(len(stillPrizes))])

	roomExchangedPrize := model.NewRoomExchangedPrize(prizeNumber, room.ID())
	if err := r.roomExchangedPrizeRepo.Save(ctx, roomExchangedPrize); err != nil {
		return nil, err
	}
	roomUser.IsExchanged = true
	if err := r.roomUserRepo.Save(ctx, roomUser); err != nil {
		return nil, err
	}

	// 他ユーザーがビンゴしたかどうか
	roomUsers, err := r.roomUserRepo.FindByRoomID(ctx, room.ID())
	if err != nil {
		return nil, err
	}
	roomBingoNumbers, err := r.roomBingoNumberRepo.FindByRoomID(ctx, room.ID())
	if err != nil {
		return nil, err
	}
	bingoNumbers := lo.Map(roomBingoNumbers, func(roomBingoNumber *model.RoomBingoNumber, _ int) uint {
		return roomBingoNumber.ID()
	})

	var bingoUsers []*model.RoomUser
	for _, roomUser := range roomUsers {
		if roomUser.IsExchanged {
			continue
		}
		bingoCard := roomUser.BingoCard(bingoNumbers)
		if bingoCard.CheckBingo() {
			bingoUsers = append(bingoUsers, roomUser)
		}
	}

	return &RequestPrizeSpinOutput{
		PrizeNumber: prizeNumber,
		UserId:      input.UserId,
		UserName:    roomUser.Name(),
		BingoUsers:  bingoUsers,
	}, nil
}
