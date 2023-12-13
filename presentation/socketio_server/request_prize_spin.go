package socketio_server

import (
	"context"
	"github.com/KaguraGateway/bingogame2-backend/application"
	socketio "github.com/googollee/go-socket.io"
	"github.com/samber/do"
)

type RequestPrizeSpin struct {
	useCase        application.RequestPrizeSpin
	socketIoServer *socketio.Server
}

func NewRequestPrizeSpin(i *do.Injector) RequestPrizeSpin {
	return RequestPrizeSpin{
		useCase:        do.MustInvoke[application.RequestPrizeSpin](i),
		socketIoServer: do.MustInvoke[*socketio.Server](i),
	}
}

func (c RequestPrizeSpin) OnEvent(conn socketio.Conn) {
	ctx := context.Background()
	output, err := c.useCase.Execute(ctx, application.RequestPrizeSpinInput{
		RoomId: conn.Rooms()[0],
		UserId: conn.Context().(string),
	})
	if err != nil {
		conn.Emit("Error", err)
		return
	}

	c.socketIoServer.ForEach("/", conn.Rooms()[0], func(conn socketio.Conn) {
		// 現在管理者とユーザーの区別つかないので、管理メッセージもまとめて流す
		conn.Emit("PrizeSpinResult", map[string]interface{}{
			"prizeNumber": output.PrizeNumber,
			"userId":      output.UserId,
			"userName":    output.UserName,
		})
	})

	// ビンゴした１人目に景品抽選許可
	if len(output.BingoUsers) > 0 {
		c.socketIoServer.ForEach("/", conn.Rooms()[0], func(conn socketio.Conn) {
			conn.Emit("UserBingo", map[string]interface{}{
				"userId":   output.BingoUsers[0].ID(),
				"userName": output.BingoUsers[0].Name(),
			})
			if output.BingoUsers[0].ID() == conn.Context().(SocketIOContext).UserId {
				conn.Emit("YourPrizeSpinTurn")
			}
		})
	}
}
