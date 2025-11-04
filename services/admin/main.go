package main

import (
	"context"
	"embed"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/pulap/pulap/pkg/lib/core"
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

	router := core.NewWebRouterWithOptions("/", core.StackOptions{
		Timeout:     60 * time.Second,
		DebugRoutes: cfg.Debug.Routes,
	}, xparams)

	var deps []any

	fileServer := core.NewFileServer(assetsFS, logger)
	deps = append(deps, fileServer)

	tmplMgr := core.NewTemplateManager(assetsFS, logger)
	deps = append(deps, tmplMgr)

	authnClient := core.NewServiceClient(cfg.Services.AuthnURL)
	userRepo := admin.NewAPIUserRepo(authnClient)
	roleRepo := admin.NewFakeRoleRepo()
	grantRepo := admin.NewFakeGrantRepo(userRepo, roleRepo)

	estateClient := core.NewServiceClient(cfg.Services.EstateURL)
	propertyRepo := admin.NewAPIPropertyRepo(estateClient)

	repos := admin.Repos{
		UserRepo:     userRepo,
		RoleRepo:     roleRepo,
		GrantRepo:    grantRepo,
		PropertyRepo: propertyRepo,
	}

	locationProvider := configureLocationProvider(cfg)
	if locationProvider == nil {
		logger.Infof("location provider disabled")
	} else {
		logger.Infof("location provider enabled: %s", locationProvider.ProviderID())
	}

	adminService := admin.NewDefaultService(repos, locationProvider, xparams)
	deps = append(deps, adminService)

	authZClient := core.NewAuthZHTTPClient(cfg.Services.AuthzURL)
	deps = append(deps, authZClient)

	dictServiceClient := core.NewServiceClient(cfg.Services.DictionaryURL)
	dictRepo := admin.NewAPIDictionaryRepo(dictServiceClient)

	adminHandler := admin.NewHandler(tmplMgr, adminService, authZClient, authnClient, dictRepo, xparams)
	deps = append(deps, adminHandler)

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

func configureLocationProvider(cfg *config.Config) admin.LocationProvider {
	if cfg == nil {
		return nil
	}
	provider := strings.ToLower(strings.TrimSpace(cfg.Geocode.Provider))
	switch provider {
	case "", admin.ProviderLocationIQ:
		key := strings.TrimSpace(cfg.Geocode.LocationIQ.Key)
		if key == "" {
			return nil
		}
		return admin.NewLocationIQProvider(admin.LocationIQOptions{
			APIKey:   key,
			Endpoint: cfg.Geocode.LocationIQ.Endpoint,
		})
	case admin.ProviderGoogle:
		return admin.NewGoogleMapsProvider(admin.GoogleMapsOptions{
			APIKey:   strings.TrimSpace(cfg.Geocode.Google.APIKey),
			Endpoint: cfg.Geocode.Google.Endpoint,
		})
	case admin.ProviderOSM:
		return admin.NewOpenStreetMapProvider(admin.OpenStreetMapOptions{
			Endpoint: cfg.Geocode.OSM.Endpoint,
			Email:    cfg.Geocode.OSM.Email,
		})
	default:
		return nil
	}
}
