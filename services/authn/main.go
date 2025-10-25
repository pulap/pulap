package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pulap/pulap/pkg/lib/core"
	"github.com/pulap/pulap/services/authn/internal/authn"
	"github.com/pulap/pulap/services/authn/internal/config"
	"github.com/pulap/pulap/services/authn/internal/mongo"
)

const (
	name    = "authn"
	version = "0.1.0"
)

func main() {
	cfg, err := config.LoadConfig("config.yaml", "AUTHN_", os.Args)
	if err != nil {
		log.Fatalf("Cannot setup %s(%s): %v", name, version, err)
	}

	logger := core.NewLogger(cfg.Log.Level)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	xparams := config.NewXParams(logger, cfg)

	corsOpts := core.DefaultCORSOptions()
	corsOpts.AllowCredentials = true
	router := core.NewRouterWithOptions(core.StackOptions{
		Timeout:     60 * time.Second,
		CORS:        &corsOpts,
		DebugRoutes: cfg.Debug.Routes,
	}, xparams)

	var deps []any

	userRepo := mongo.NewUserMongoRepo(xparams)
	deps = append(deps, userRepo)

	UserHandler := authn.NewUserHandler(userRepo, xparams)
	deps = append(deps, UserHandler)

	AuthHandler := authn.NewAuthHandler(userRepo, xparams)
	deps = append(deps, AuthHandler)

	systemHandler := authn.NewSystemHandler(userRepo, xparams)
	deps = append(deps, systemHandler)

	starts, stops, _ := core.Setup(ctx, router, deps...)

	if err := core.Start(ctx, starts, stops); err != nil {
		logger.Errorf("Cannot start %s(%s): %v", name, version, err)
		os.Exit(1)
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
