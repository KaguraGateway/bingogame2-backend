package socketio_server

import (
	"context"
	"encoding/json"
	"fmt"
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
	fmt.Printf("%v\n", msg)
	var input RoomUserRegisterInput
	if err := json.Unmarshal([]byte(msg), &input); err != nil {
		conn.Emit("Error", err)
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("%v\n", conn.Rooms())

	ctx := context.Background()
	output, err := c.useCase.Execute(ctx, application.RoomUserRegisterInput{
		RoomId:   conn.Rooms()[0],
		UserName: input.UserName,
	})
	if err != nil {
		conn.Emit("Error", err)
		fmt.Printf("Error: %v\n", err)
		return
	}

	conn.SetContext(SocketIOContext{
		UserId: output.UserId,
	})

	fmt.Printf("%v\n", output)

	conn.Emit("UserAuthSuccess", output)
}
