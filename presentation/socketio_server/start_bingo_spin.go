package socketio_server

import (
	"context"
	"github.com/KaguraGateway/bingogame2-backend/application"
	"github.com/KaguraGateway/bingogame2-backend/domain/model"
	socketio "github.com/googollee/go-socket.io"
	"github.com/samber/do"
	"github.com/samber/lo"
)

type StartBingoSpin struct {
	useCase        application.StartBingoSpin
	socketIoServer *socketio.Server
}

func NewStartBingoSpin(i *do.Injector) StartBingoSpin {
	return StartBingoSpin{
		useCase:        do.MustInvoke[application.StartBingoSpin](i),
		socketIoServer: do.MustInvoke[*socketio.Server](i),
	}
}

func (c StartBingoSpin) OnEvent(conn socketio.Conn) {
	ctx := context.Background()
	output, err := c.useCase.Execute(ctx, application.StartBingoSpinInput{
		RoomId: conn.Rooms()[0],
	})
	if err != nil {
		conn.Emit("Error", err)
		return
	}

	c.socketIoServer.ForEach("/", conn.Rooms()[0], func(conn socketio.Conn) {
		// 現在管理者とユーザーの区別つかないので、管理メッセージもまとめて流す
		conn.Emit("BingoSpinResult", map[string]interface{}{
			"bingoNumber": output.BingoNumber,
		})
		// ユーザーのビンゴカード更新
		conn.Emit("UpdateUserBingoCard", map[string]interface{}{
			"bingoCard": output.BingoCards[conn.Context().(SocketIOContext).UserId],
		})
		// ビンゴしてたら通知
		if lo.ContainsBy(output.BingoUsers, func(item *model.RoomUser) bool {
			return item.ID() == conn.Context().(SocketIOContext).UserId
		}) {
			conn.Emit("Bingo")
		}
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
