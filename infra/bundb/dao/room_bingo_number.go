package dao

type RoomBingoNumber struct {
	ID     uint   `bun:",pk"`
	RoomID string `bun:",pk"`
}
