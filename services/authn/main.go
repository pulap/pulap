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
		Timeout: 60 * time.Second,
		CORS:    &corsOpts,
	}, xparams)

	var deps []any

	// WIP: Temporary repository selector.
	// A single implementation will remain later.
	var userRepo authn.UserRepo
	if cfg.Database.MongoURL != "" {
		userRepo = mongo.NewUserMongoRepo(xparams)
		logger.Infof("Using MongoDB repository: %s", cfg.Database.MongoURL)
	} else {
		// TODO: Add SQLite repo when needed
		userRepo = mongo.NewUserMongoRepo(xparams)
		logger.Infof("SQLite not implemented yet, falling back to MongoDB")
	}
	deps = append(deps, userRepo)

	UserHandler := authn.NewUserHandler(userRepo, xparams)
	deps = append(deps, UserHandler)

	AuthHandler := authn.NewAuthHandler(userRepo, xparams)
	deps = append(deps, AuthHandler)

	// Register system handler so bootstrap endpoints are exposed
	systemHandler := authn.NewSystemHandler(userRepo, xparams)
	deps = append(deps, systemHandler)

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
