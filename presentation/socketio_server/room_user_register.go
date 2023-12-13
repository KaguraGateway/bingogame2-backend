package socketio_server

import (
	"context"
	"encoding/json"
	"github.com/KaguraGateway/bingogame2-backend/application"
	socketio "github.com/googollee/go-socket.io"
	"github.com/samber/do"
)

type RoomUserRegister struct {
	useCase application.RoomUserRegister
}

func NewRoomUserRegister(i *do.Injector) RoomUserRegister {
	return RoomUserRegister{
		useCase: do.MustInvoke[application.RoomUserRegister](i),
	}
}

type RoomUserRegisterInput struct {
	UserName string `json:"userName"`
}

func (c RoomUserRegister) OnEvent(conn socketio.Conn, msg string) {
	var input RoomUserRegisterInput
	if err := json.Unmarshal([]byte(msg), &input); err != nil {
		conn.Emit("Error", err)
		return
	}

	ctx := context.Background()
	output, err := c.useCase.Execute(ctx, application.RoomUserRegisterInput{
		RoomId:   conn.Rooms()[0],
		UserName: input.UserName,
	})
	if err != nil {
		conn.Emit("Error", err)
		return
	}

	conn.SetContext(SocketIOContext{
		UserId: output.UserId,
	})

	conn.Emit("UserAuthSuccess", output)
}
