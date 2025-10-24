package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"

	"github.com/pulap/pulap/pkg/lib/core"
	"github.com/pulap/pulap/pkg/lib/core/middleware"
	"github.com/pulap/pulap/services/authz/internal/authz"
	"github.com/pulap/pulap/services/authz/internal/config"
	"github.com/pulap/pulap/services/authz/internal/mongo"
)

const (
	name    = "authz"
	version = "0.1.0"
)

func main() {
	cfg, err := config.LoadConfig("config.yaml", "AUTHZ_", os.Args)
	if err != nil {
		log.Fatalf("Cannot setup %s(%s): %v", name, version, err)
	}

	logger := core.NewLogger(cfg.Log.Level)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	xparams := config.XParams{
		Log: logger,
		Cfg: cfg,
	}

	router := chi.NewRouter()
	router.Use(middleware.RequestIDMiddleware)

	var deps []any

	// WIP: Temporary repository selector.
	// A single implementation will remain later.
	var roleRepo authz.RoleRepo
	var grantRepo authz.GrantRepo

	if cfg.Database.MongoURL != "" {
		roleRepo = mongo.NewRoleMongoRepo(xparams)
		grantRepo = mongo.NewGrantMongoRepo(xparams)
		logger.Infof("Using MongoDB repository: %s", cfg.Database.MongoURL)
	} else {
		// TODO: Add SQLite repo when needed
		roleRepo = mongo.NewRoleMongoRepo(xparams)
		grantRepo = mongo.NewGrantMongoRepo(xparams)
		logger.Infof("SQLite not implemented yet, falling back to MongoDB")
	}

	deps = append(deps, roleRepo)
	deps = append(deps, grantRepo)

	// Policy engine setup
	policyEngine := authz.NewPolicyEngine(roleRepo, grantRepo)

	// Handler setup
	roleHandler := authz.NewRoleHandler(roleRepo, xparams)
	deps = append(deps, roleHandler)

	grantHandler := authz.NewGrantHandler(grantRepo, roleRepo, xparams)
	deps = append(deps, grantHandler)

	policyHandler := authz.NewPolicyHandler(policyEngine, xparams)
	deps = append(deps, policyHandler)

	starts, stops := core.Setup(ctx, router, deps...)

	if err := core.Start(ctx, starts, stops); err != nil {
		logger.Errorf("Cannot start %s(%s): %v", name, version, err)
		log.Fatal(err)
	}

	// Bootstrap service setup
	bootstrapService := authz.NewBootstrapService(roleRepo, grantRepo, xparams)

	// Run bootstrap process (log but don't fail startup)
	if err := bootstrapService.Bootstrap(ctx); err != nil {
		logger.Errorf("Bootstrap failed: %v", err)
		// don't fail startup
	} else {
		logger.Infof("Bootstrap completed successfully")
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
