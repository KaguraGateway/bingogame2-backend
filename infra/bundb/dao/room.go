package dao

type Room struct {
	ID       string `bun:",pk"`
	PrizeNum uint   `bun:",notnull"`
}
