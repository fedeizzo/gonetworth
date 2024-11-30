package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/rs/zerolog"

	"github.com/fedeizzo/gonetworth/internal/app"
	"github.com/fedeizzo/gonetworth/internal/config"
	"github.com/fedeizzo/gonetworth/internal/db"
	"github.com/fedeizzo/gonetworth/internal/web"
)

func main() {
	ctx := context.Background()
	cfg := config.Parse()
	logger := initLogger(&cfg)
	ctx = logger.WithContext(ctx)

	dbURL, err := cfg.DatabaseUrl()
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("cannot retrieve the database URL from the config")

		os.Exit(1)
	}

	conn, err := pgx.Connect(context.Background(), dbURL.String())
	if err != nil {
		dbURL, _ := cfg.DatabaseUrl()
		logger.Fatal().
			Err(err).
			Str("database_url", dbURL.String()).
			Msg("cannot connect to the database")

		os.Exit(1)
	}

	querier := db.New(conn)
	app := app.App{Querier: querier}

	logger.Info().
		Msg("connected to the db")

	accountService := app.GetAccountService()
	accounts, err := accountService.GetExpenseAccounts(ctx)
	if err != nil {
		panic(err)
	}

	var names []string
	for _, a := range accounts {
		names = append(names, a.Name)
	}
	fmt.Println(fuzzy.Find("auc", names))

	if err := web.Run(); err != nil {
		logger.Fatal().
			Err(err).
			Msg("cannot run the server")

		os.Exit(1)
	}
}

func initLogger(cfg *config.Config) zerolog.Logger {
	var out io.Writer = os.Stdout
	if cfg.LogFormat == config.FormatText {
		out = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}
	}

	return zerolog.New(out).Level(cfg.LogLevel).With().Timestamp().Caller().Logger()
}
