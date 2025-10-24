package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/pulap/pulap/pkg/lib/core"
	coremw "github.com/pulap/pulap/pkg/lib/core/middleware"
	"github.com/pulap/pulap/services/estate/internal/config"
	"github.com/pulap/pulap/services/estate/internal/estate"
	"github.com/pulap/pulap/services/estate/internal/sqlite"
)

const (
	name    = "estate"
	version = "0.1.0"
)

func main() {
	cfg, err := config.LoadConfig("config.yaml", "APP", os.Args)
	if err != nil {
		log.Fatalf("Cannot setup %s(%s): %v", name, version, err)
	}

	logger := core.NewLogger(cfg.Log.Level)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	xparams := config.NewXParams(logger, cfg)

	router := chi.NewRouter()
	corsOpts := coremw.DefaultCORSOptions()
	corsOpts.AllowCredentials = true
	coremw.ApplyStack(router, logger, coremw.StackOptions{
		Timeout: 60 * time.Second,
		CORS:    &corsOpts,
	})

	var deps []any
	EstateRepo := sqlite.NewEstateSQLiteRepo(xparams)
	deps = append(deps, EstateRepo)

	EstateHandler := estate.NewEstateHandler(EstateRepo, xparams)
	deps = append(deps, EstateHandler)

	starts, stops := core.Setup(ctx, router, deps...)

	if err := core.Start(ctx, starts, stops); err != nil {
		logger.Errorf("Cannot start %s(%s): %v", name, version, err)
		log.Fatal(err)
	}

	logger.Infof("%s(%s) started successfully", name, version)

	go func() {
		core.Serve(router, core.ServerOpts{Port: cfg.Server.Port}, stops, logger)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	<-stop

	logger.Infof("Shutting down %s(%s)...", name, version)
	cancel()
}
