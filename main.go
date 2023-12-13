package main

import (
	"database/sql"
	"fmt"
	"github.com/KaguraGateway/bingogame2-backend/application"
	"github.com/KaguraGateway/bingogame2-backend/infra/bundb"
	"github.com/KaguraGateway/bingogame2-backend/presentation/http_server"
	"github.com/KaguraGateway/bingogame2-backend/presentation/socketio_server"
	"github.com/getsentry/sentry-go"
	socketio "github.com/googollee/go-socket.io"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/samber/do"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"log"
	"os"
	"strconv"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Print(err)
	}

	// Sentry
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:           os.Getenv("SENTRY_DSN"),
		EnableTracing: true,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v", err)
	}

	// Start DB
	sqlDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv("DATABASE_URL"))))
	db := bun.NewDB(sqlDb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	defer db.Close()

	// Start server
	// SocketIO
	socketIoServer := socketio.NewServer(nil)
	socketio_server.RegisterRoutes(socketIoServer, i)
	go socketIoServer.Serve()
	defer socketIoServer.Close()

	// Build Injector
	i := buildInjector(db, socketIoServer)

	// HTTP Server
	httpServer := echo.New()
	httpServer.HideBanner = true
	http_server.RegisterRoutes(httpServer, socketIoServer, i)

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}

	httpServer.Logger.Fatal(httpServer.Start(fmt.Sprintf(":%d", port)))
}

func buildInjector(db *bun.DB, socketIoServer *socketio.Server) *do.Injector {
	i := do.New()

	// Register DB
	do.Provide(i, func(i *do.Injector) (*bun.DB, error) {
		return db, nil
	})
	// Register SocketIO
	do.Provide(i, func(i *do.Injector) (*socketio.Server, error) {
		return socketIoServer, nil
	})
	// Register Repository
	do.Provide(i, bundb.NewRoomRepository)
	do.Provide(i, bundb.NewRoomBingoNumberRepository)
	do.Provide(i, bundb.NewRoomExchangedPrizeRepository)
	do.Provide(i, bundb.NewRoomUserRepository)
	// Register UseCase
	do.Provide(i, application.NewCreateRoomUseCase)
	do.Provide(i, application.NewRequestAdminInitUseCase)
	do.Provide(i, application.NewRequestPrizeSpinUseCase)
	do.Provide(i, application.NewRoomUserAuthUseCase)
	do.Provide(i, application.NewRoomUserRegisterUseCase)
	do.Provide(i, application.NewStartBingoSpinUseCase)

	return i
}
