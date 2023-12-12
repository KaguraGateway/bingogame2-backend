package model

import "github.com/oklog/ulid/v2"

type Room struct {
	id       string
	prizeNum uint
}

func NewRoom(prizeNum uint) *Room {
	return &Room{
		id:       ulid.Make().String(),
		prizeNum: prizeNum,
	}
}

func RebuildRoom(id string, prizeNum uint) *Room {
	return &Room{
		id:       id,
		prizeNum: prizeNum,
	}
}

func (r *Room) ID() string {
	return r.id
}

func (r *Room) PrizeNum() uint {
	return r.prizeNum
}
