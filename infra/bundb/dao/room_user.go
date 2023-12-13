package dao

type RoomUser struct {
	ID          string `bun:",pk"`
	RoomID      string `bun:",notnull"`
	Name        string `bun:",notnull"`
	BingoSeed   int64  `bun:",notnull"`
	IsExchanged bool   `bun:",notnull"`
}
