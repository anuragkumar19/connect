package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/anuragkumar19/connect/api"
	"github.com/anuragkumar19/connect/database"
	"github.com/anuragkumar19/connect/infra/nats"
	"github.com/anuragkumar19/connect/infra/postgres"
	"github.com/anuragkumar19/connect/infra/smtp"
	"github.com/anuragkumar19/connect/infra/storage"
	"github.com/anuragkumar19/connect/mailer"
	"github.com/anuragkumar19/connect/server"
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

	ctxBase, cancelBase := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelBase()
	ctx, cancel := signal.NotifyContext(ctxBase, syscall.SIGINT, syscall.SIGTERM)
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

	smtpCfg, err := smtp.AutoConfig()
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to parse smtp config")
	}
	smtpClient, err := smtp.New(ctx, smtpCfg, &logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("smtp init failed")
	}

	// TODO: graceful shutdown
	// TODO: background services and consumers
	// TODO: prometheus

	database := database.New(pgConn)

	api := api.New(&logger, database, &natsConn, &storageClient, &smtpClient)

	serverCfg, err := server.AutoConfig()
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to parse server config")
	}
	server, err := server.New(serverCfg, &logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to create server")
	}

	if err := server.Mount("/api", api.Router()); err != nil {
		logger.Fatal().Err(err).Msg("failed to mount api router to server")
	}

	go func() {
		if err := server.Start(); err != nil {
			logger.Fatal().Err(err).Msg("")
		}
	}()

	mailerServerCfg, err := mailer.AutoConfig()
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to parse mailerServer config")
	}
	mailerServer, err := mailer.NewServer(ctx, mailerServerCfg, &logger, &smtpClient, &natsConn)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to create mailer server")
	}
	go func() {
		if err := mailerServer.Start(); err != nil {
			logger.Fatal().Err(err).Msg("")
		}
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	sig := <-signalCh
	logger.Info().Str("signal", sig.String()).Msg("signal received shutting down servers")

	shutDownCtxBase, shutdownCancelBase := context.WithTimeout(context.Background(), 20*time.Second)
	defer shutdownCancelBase()
	shutDownCtx, shutdownCancel := signal.NotifyContext(shutDownCtxBase, syscall.SIGINT, syscall.SIGTERM)
	defer shutdownCancel()
	if err := server.Shutdown(shutDownCtx); err != nil {
		logger.Fatal().Err(err).Msg("failed to shutdown server")
	}
	if err := mailerServer.Shutdown(shutDownCtx); err != nil {
		logger.Fatal().Err(err).Msg("failed to mailer shutdown server")
	}
}
