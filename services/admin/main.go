package main

import (
	"context"
	"embed"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/pulap/pulap/pkg/lib/core"
	coremw "github.com/pulap/pulap/pkg/lib/core/middleware"
	"github.com/pulap/pulap/services/admin/internal/admin"
	"github.com/pulap/pulap/services/admin/internal/config"
)

const (
	name    = "admin"
	version = "0.1.0"
)

//go:embed assets
var assetsFS embed.FS

func main() {
	cfg, err := config.LoadConfig("config.yaml", "ADMIN", os.Args)
	if err != nil {
		log.Fatalf("Cannot setup %s(%s): %v", name, version, err)
	}

	logger := core.NewLogger(cfg.Log.Level)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	xparams := config.NewXParams(logger, cfg)

	router := chi.NewRouter()
	coremw.ApplyStack(router, logger, coremw.StackOptions{Timeout: 60 * time.Second})
	router.Use(chimiddleware.NoCache)

	var deps []any

	fileServer := core.NewFileServer(assetsFS, logger)
	deps = append(deps, fileServer)

	tmplMgr := core.NewTemplateManager(assetsFS, logger)
	deps = append(deps, tmplMgr)

	authnClient := core.NewServiceClient(cfg.Services.AuthnURL)
	userRepo := admin.NewAPIUserRepo(authnClient)
	roleRepo := admin.NewFakeRoleRepo()
	grantRepo := admin.NewFakeGrantRepo(userRepo, roleRepo)

	repos := admin.Repos{
		UserRepo:  userRepo,
		RoleRepo:  roleRepo,
		GrantRepo: grantRepo,
	}

	adminService := admin.NewDefaultService(repos, xparams)
	deps = append(deps, adminService)

	authZClient := core.NewAuthZHTTPClient(cfg.Services.AuthzURL)
	deps = append(deps, authZClient)

	adminHandler := admin.NewHandler(tmplMgr, adminService, authZClient, authnClient, xparams)
	deps = append(deps, adminHandler)

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
