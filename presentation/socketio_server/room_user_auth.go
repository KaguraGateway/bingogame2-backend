package socketio_server

import (
	"context"
	"encoding/json"
	"github.com/KaguraGateway/bingogame2-backend/application"
	socketio "github.com/googollee/go-socket.io"
	"github.com/samber/do"
)

type RoomUserAuth struct {
	useCase application.RoomUserAuth
}

func NewRoomUserAuth(i *do.Injector) RoomUserAuth {
	return RoomUserAuth{
		useCase: do.MustInvoke[application.RoomUserAuth](i),
	}
}

type RoomUserAuthInput struct {
	UserId string `json:"userId"`
}

func (c RoomUserAuth) OnEvent(conn socketio.Conn, msg string) {
	var input RoomUserAuthInput
	if err := json.Unmarshal([]byte(msg), &input); err != nil {
		conn.Emit("Error", err)
		return
	}

	ctx := context.Background()
	output, err := c.useCase.Execute(ctx, application.RoomUserAuthInput{
		RoomId: conn.Rooms()[0],
		UserId: input.UserId,
	})
	if err != nil {
		conn.Emit("Error", err)
		return
	}

	conn.SetContext(SocketIOContext{
		UserId: input.UserId,
	})

	conn.Emit("UserAuthSuccess", output)
}
