package socketio_server

import (
	"context"
	"github.com/KaguraGateway/bingogame2-backend/application"
	socketio "github.com/googollee/go-socket.io"
	"github.com/samber/do"
)

type RequestAdminInit struct {
	useCase application.RequestAdminInit
}

func NewRequestAdminInit(i *do.Injector) RequestAdminInit {
	return RequestAdminInit{
		useCase: do.MustInvoke[application.RequestAdminInit](i),
	}
}

func (c RequestAdminInit) OnEvent(conn socketio.Conn) {
	ctx := context.Background()
	output, err := c.useCase.Execute(ctx, application.RequestAdminInitInput{
		RoomId: conn.Rooms()[0],
	})
	if err != nil {
		conn.Emit("Error", err)
		return
	}

	conn.Emit("RequestAdminInitSuccess", output)
}
