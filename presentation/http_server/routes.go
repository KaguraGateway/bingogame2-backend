package http_server

import (
	socketio "github.com/googollee/go-socket.io"
	"github.com/labstack/echo/v4"
	"github.com/samber/do"
)

func RegisterRoutes(app *echo.Echo, socketIoServer *socketio.Server, i *do.Injector) {
	app.Any("/socket.io/", func(ctx echo.Context) error {
		socketIoServer.ServeHTTP(ctx.Response(), ctx.Request())
		return nil
	})
	app.POST("/room/create", NewCreateRoom(i).POST)
}
