package main

import (
	"database/sql"
	"fmt"
	"github.com/KaguraGateway/bingogame2-backend/cmd/bin/migrations"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/migrate"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	app := &cli.App{
		Name: "migrations",
		Commands: []*cli.Command{
			dbCommand(),
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func db() *bun.DB {
	sqlDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv("DATABASE_URL"))))
	return bun.NewDB(sqlDb, pgdialect.New())
}

func dbCommand() *cli.Command {
	return &cli.Command{
		Name: "db",
		Subcommands: []*cli.Command{
			{
				Name: "init",
				Action: func(ctx *cli.Context) error {
					db := db()
					defer db.Close()
					migrator := migrate.NewMigrator(db, migrations.Migrations)
					return migrator.Init(ctx.Context)
				},
			},
			{
				Name: "migrate",
				Action: func(ctx *cli.Context) error {
					db := db()
					defer db.Close()
					migrator := migrate.NewMigrator(db, migrations.Migrations)
					group, err := migrator.Migrate(ctx.Context)
					if err != nil {
						return err
					}
					if group.ID == 0 {
						fmt.Println("no changes")
						return nil
					}
					fmt.Printf("migrated to %s\n", group)
					return nil
				},
			},
			{
				Name: "create",
				Action: func(ctx *cli.Context) error {
					db := db()
					defer db.Close()
					migrator := migrate.NewMigrator(db, migrations.Migrations)
					name := strings.Join(ctx.Args().Slice(), "_")
					mf, err := migrator.CreateGoMigration(ctx.Context, name)
					if err != nil {
						return err
					}
					fmt.Printf("created %s\n", mf)
					return nil
				},
			},
		},
	}
}
