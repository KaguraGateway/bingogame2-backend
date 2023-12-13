package socketio_server

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"github.com/samber/do"
)

type SocketIOContext struct {
	UserId string
}

func RegisterRoutes(socketIoServer *socketio.Server, i *do.Injector) {
	socketIoServer.OnConnect("/", func(conn socketio.Conn) error {
		url := conn.URL()
		roomId := url.Query().Get("room_id")
		fmt.Printf("NewConnect: %v\n", roomId)

		conn.Join(roomId)

		return nil
	})
	// User
	socketIoServer.OnEvent("/", "UserAuth", NewRoomUserAuth(i).OnEvent)
	socketIoServer.OnEvent("/", "UserRegister", NewRoomUserAuth(i).OnEvent)
	socketIoServer.OnEvent("/", "RequestPrizeSpin", NewRequestPrizeSpin(i).OnEvent)
	// Admin
	socketIoServer.OnEvent("/", "RequestAdminInit", NewRequestAdminInit(i).OnEvent)
	socketIoServer.OnEvent("/", "StartBingoSpin", NewStartBingoSpin(i).OnEvent)
}
