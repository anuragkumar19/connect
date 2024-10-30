package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/anuragkumar19/connect/api"
	"github.com/anuragkumar19/connect/database"
	"github.com/anuragkumar19/connect/infra/nats"
	"github.com/anuragkumar19/connect/infra/postgres"
	"github.com/anuragkumar19/connect/infra/storage"
	"github.com/rs/zerolog"
	"go.uber.org/automaxprocs/maxprocs"
)

var (
	version           = "0.0.1"
	commit            = "NA"
	buildBranch       = "main"
	buildStampRFC3339 string
	buildStamp        time.Time
)

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("bin", "connect-server").Logger()

	logLevel := zerolog.InfoLevel

	if v := os.Getenv("LOG_LEVEL"); v != "" {
		lvl, err := zerolog.ParseLevel(v)
		if err != nil {
			logger.Fatal().Err(err).Msg("cannot parse env variable LOG_LEVEL")
		}
		logLevel = lvl
	}
	logger = logger.Level(logLevel)

	buildStamp = time.Now()

	if buildStampRFC3339 != "" {
		bs, err := time.Parse(time.RFC3339, buildStampRFC3339)
		if err != nil {
			logger.Fatal().Err(err).Msg("failed to parse main.buildStampRFC3339 which is linked during build")
		}
		buildStamp = bs
	}

	logger.Info().Str("version", version).Str("commit", commit).Str("branch", buildBranch).Time("build_at", buildStamp).Msg("")

	maxprocs.Set(maxprocs.Logger(func(s string, i ...interface{}) {
		str := fmt.Sprintf(s, i...)
		logger.Info().Msg(str)
	}))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	pgConfig, err := postgres.AutoConfig()
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to parse postgres config")
	}
	pgConn, err := postgres.New(ctx, pgConfig, &logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("postgres init failed")
	}

	natsCfg, err := nats.AutoConfig()
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to parse nats config")
	}
	natsConn, err := nats.New(ctx, natsCfg, &logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("nats init failed")
	}

	storageCfg, err := storage.AutoConfig()
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to parse storage config")
	}
	storageClient, err := storage.New(ctx, storageCfg, &logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("storage init failed")
	}

	// TODO: graceful shutdown
	// TODO: background services and consumers
	// TODO: prometheus

	database := database.New(pgConn)

	server := api.NewServer(&logger, database, &natsConn, &storageClient)

	if err := server.Serve(); err != nil {
		logger.Fatal().Err(err).Send()
	}
}
