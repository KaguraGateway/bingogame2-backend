package model

type RoomExchangedPrize struct {
	id     uint
	roomID string
}

func NewRoomExchangedPrize(id uint, roomID string) *RoomExchangedPrize {
	return &RoomExchangedPrize{
		id:     id,
		roomID: roomID,
	}
}

func RebuildRoomExchangedPrize(id uint, roomID string) *RoomExchangedPrize {
	return &RoomExchangedPrize{
		id:     id,
		roomID: roomID,
	}
}

func (rep *RoomExchangedPrize) ID() uint {
	return rep.id
}

func (rep *RoomExchangedPrize) RoomID() string {
	return rep.roomID
}
