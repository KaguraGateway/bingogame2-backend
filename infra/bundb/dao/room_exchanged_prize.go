package dao

type RoomExchangedPrize struct {
	ID     uint   `bun:",pk"`
	RoomID string `bun:",pk"`
}
