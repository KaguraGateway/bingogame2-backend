package model

type RoomBingoNumber struct {
	id     uint
	roomID string
}

func NewRoomBingoNumber(id uint, roomID string) *RoomBingoNumber {
	return &RoomBingoNumber{
		id:     id,
		roomID: roomID,
	}
}

func RebuildRoomBingoNumber(id uint, roomID string) *RoomBingoNumber {
	return &RoomBingoNumber{
		id:     id,
		roomID: roomID,
	}
}

func (rbn *RoomBingoNumber) ID() uint {
	return rbn.id
}

func (rbn *RoomBingoNumber) RoomID() string {
	return rbn.roomID
}
