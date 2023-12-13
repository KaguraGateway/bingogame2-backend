package http_server

import (
	"github.com/KaguraGateway/bingogame2-backend/application"
	"github.com/labstack/echo/v4"
	"github.com/samber/do"
)

type CreateRoom struct {
	createRoomUseCase application.CreateRoom
}

func NewCreateRoom(i *do.Injector) CreateRoom {
	return CreateRoom{
		createRoomUseCase: do.MustInvoke[application.CreateRoom](i),
	}
}

type CreateRoomInput struct {
	PrizeNum uint `json:"prizeNum"`
}

func (c CreateRoom) POST(ctx echo.Context) error {
	var input CreateRoomInput
	if err := ctx.Bind(&input); err != nil {
		return err
	}
	output, err := c.createRoomUseCase.Execute(ctx.Request().Context(), application.CreateRoomInput{
		PrizeNum: input.PrizeNum,
	})
	if err != nil {
		return err
	}
	return ctx.JSON(200, output)
}
