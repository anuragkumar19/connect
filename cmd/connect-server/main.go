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
	"github.com/anuragkumar19/connect/metrics"
	"github.com/anuragkumar19/connect/pkg/buildinfo"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.uber.org/automaxprocs/maxprocs"
)

var (
	bin               = "connect-server"
	version           = "0.0.1"
	commit            = "NA"
	buildBranch       = "main"
	buildStampRFC3339 string
	buildStamp        time.Time
)

func main() {
	// TODO: Register hook for tracing. Use otel-slog in different codebase and see results read the third party bridge code and impl own hook
	log.Logger = zerolog.New(os.Stdout).With().Caller().Timestamp().Logger()

	logLevel := zerolog.InfoLevel

	if v := os.Getenv("LOG_LEVEL"); v != "" {
		lvl, err := zerolog.ParseLevel(v)
		if err != nil {
			log.Fatal().Err(err).Msg("cannot parse env variable LOG_LEVEL")
		}
		logLevel = lvl
	}
	log.Logger = log.Logger.Level(logLevel)

	buildStamp = time.Now()

	if buildStampRFC3339 != "" {
		bs, err := time.Parse(time.RFC3339, buildStampRFC3339)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to parse main.buildStampRFC3339 which is linked during build")
		}
		buildStamp = bs
	}

	buildinfo.SetBin(bin)
	buildinfo.SetVersion(version)
	buildinfo.SetBranch(buildBranch)
	buildinfo.SetCommit(commit)
	buildinfo.SetTimestamp(buildStamp)

	log.Info().Str("version", version).Str("bin", bin).Str("commit", commit).Str("branch", buildBranch).Time("build_at", buildStamp).Msg("")

	maxprocs.Set(maxprocs.Logger(func(s string, i ...interface{}) {
		str := fmt.Sprintf(s, i...)
		log.Info().Msg(str)
	}))

	ctxBase, cancelBase := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelBase()
	ctx, cancel := signal.NotifyContext(ctxBase, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	pgConfig, err := postgres.AutoConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse postgres config")
	}
	pgConn, err := postgres.New(ctx, pgConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("postgres init failed")
	}

	natsCfg, err := nats.AutoConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse nats config")
	}
	natsConn, err := nats.New(ctx, natsCfg)
	if err != nil {
		log.Fatal().Err(err).Msg("nats init failed")
	}

	storageCfg, err := storage.AutoConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse storage config")
	}
	storageClient, err := storage.New(ctx, storageCfg)
	if err != nil {
		log.Fatal().Err(err).Msg("storage init failed")
	}

	smtpCfg, err := smtp.AutoConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse smtp config")
	}
	smtpClient, err := smtp.New(ctx, smtpCfg)
	if err != nil {
		log.Fatal().Err(err).Msg("smtp init failed")
	}

	database := database.New(pgConn)

	apiConf, err := api.AutoConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse api config")
	}
	apiServer, err := api.NewServer(apiConf, database, &natsConn, &storageClient, &smtpClient)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create api server")
	}

	metricsCfg, err := metrics.AutoConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse metrics config")
	}
	metricsServer, err := metrics.NewServer(metricsCfg)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create metrics server")
	}

	go func() {
		if err := apiServer.Start(); err != nil {
			log.Fatal().Err(err).Msg("")
		}
	}()

	go func() {
		if err := metricsServer.Start(); err != nil {
			log.Fatal().Err(err).Msg("")
		}
	}()

	mailerServerCfg, err := mailer.AutoConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse mailerServer config")
	}
	mailerServer, err := mailer.NewServer(ctx, mailerServerCfg, &smtpClient, &natsConn)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create mailer server")
	}
	go func() {
		if err := mailerServer.Start(); err != nil {
			log.Fatal().Err(err).Msg("")
		}
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	sig := <-signalCh
	log.Info().Str("signal", sig.String()).Msg("signal received shutting down servers")

	shutDownCtxBase, shutdownCancelBase := context.WithTimeout(context.Background(), 20*time.Second)
	defer shutdownCancelBase()
	shutDownCtx, shutdownCancel := signal.NotifyContext(shutDownCtxBase, syscall.SIGINT, syscall.SIGTERM)
	defer shutdownCancel()
	if err := apiServer.Shutdown(shutDownCtx); err != nil {
		log.Fatal().Err(err).Msg("failed to shutdown server")
	}
	if err := metricsServer.Shutdown(shutDownCtx); err != nil {
		log.Fatal().Err(err).Msg("failed to shutdown server")
	}
	if err := mailerServer.Shutdown(shutDownCtx); err != nil {
		log.Fatal().Err(err).Msg("failed to mailer shutdown server")
	}

	pgConn.Close()
	if err := natsConn.DrainAndWaitForClose(shutDownCtx); err != nil {
		log.Fatal().Err(err).Msg("failed to drain and close nats conn")
	}
}
