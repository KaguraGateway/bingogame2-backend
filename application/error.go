package application

import "github.com/cockroachdb/errors"

var ErrRoomNotFound = errors.New("room not found")
var ErrRoomUserNotFound = errors.New("room user not found")
var ErrFailedCreateRoomUser = errors.New("failed create room user")
var ErrAlreadyExchanged = errors.New("already exchanged")
