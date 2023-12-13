package model

import (
	"crypto/rand"
	"encoding/binary"
	"github.com/KaguraGateway/bingogame2-backend/domain"
	"github.com/oklog/ulid/v2"
)

type RoomUser struct {
	id          string
	roomID      string
	name        string
	bingoSeed   int64
	IsExchanged bool
}

func NewRoomUser(roomID string, name string) *RoomUser {
	var bingoSeed int64
	if err := binary.Read(rand.Reader, binary.LittleEndian, &bingoSeed); err != nil {
		panic(err)
	}

	return &RoomUser{
		id:          ulid.Make().String(),
		roomID:      roomID,
		name:        name,
		bingoSeed:   bingoSeed,
		IsExchanged: false,
	}
}

func RebuildRoomUser(id string, roomID string, name string, bingoSeed int64, isExchanged bool) *RoomUser {
	return &RoomUser{
		id:          id,
		roomID:      roomID,
		name:        name,
		bingoSeed:   bingoSeed,
		IsExchanged: isExchanged,
	}
}

func (ru *RoomUser) ID() string {
	return ru.id
}

func (ru *RoomUser) RoomID() string {
	return ru.roomID
}

func (ru *RoomUser) Name() string {
	return ru.name
}

// TODO: 事案なのでリファクタリングしたい
func (ru *RoomUser) BingoCard(bingoNumbers []uint) *BingoCard {
	return RebuildBingoCard(ru.bingoSeed, bingoNumbers)
}

func (ru *RoomUser) BingoSeed() int64 {
	return ru.bingoSeed
}

func (ru *RoomUser) SetName(name string) error {
	if len(name) == 0 {
		return domain.ErrInvalidName
	}
	ru.name = name
	return nil
}
