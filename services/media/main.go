package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pulap/pulap/pkg/lib/core"
	"github.com/pulap/pulap/services/media/internal/config"
	"github.com/pulap/pulap/services/media/internal/dictionary"
	"github.com/pulap/pulap/services/media/internal/media"
	"github.com/pulap/pulap/services/media/internal/storage"
)

const (
	name    = "media"
	version = "0.1.0"
)

func main() {
	cfg, err := config.LoadConfig("config.yaml", "MEDIA", os.Args)
	if err != nil {
		log.Fatalf("cannot setup %s(%s): %v", name, version, err)
	}

	logger := core.NewLogger(cfg.Log.Level)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	xparams := config.NewXParams(logger, cfg)

	router := core.NewRouterWithOptions(core.StackOptions{
		Timeout:     60 * time.Second,
		DebugRoutes: cfg.Debug.Routes,
	}, xparams)

	storageBackend, err := configureStorage(cfg)
	if err != nil {
		logger.Errorf("cannot configure storage backend: %v", err)
		os.Exit(1)
	}

	repo := media.NewInMemoryRepository()
	dictClient := dictionary.NewNoopClient()

	var variants []media.VariantDefinition
	for _, v := range cfg.Processing.Variants {
		variants = append(variants, media.VariantDefinition{
			Name:   v.Name,
			Width:  v.Width,
			Height: v.Height,
		})
	}

	service := media.NewService(repo, storageBackend, dictClient, media.ServiceOptions{
		EnableCropping:    cfg.Processing.Cropping,
		EnableCompression: cfg.Processing.Compression,
		Variants:          variants,
	})

	handler := media.NewHandler(service, xparams)

	starts, stops, _ := core.Setup(ctx, router, handler)

	if err := core.Start(ctx, starts, stops); err != nil {
		logger.Errorf("cannot start %s(%s): %v", name, version, err)
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

func configureStorage(cfg *config.Config) (storage.MediaStorage, error) {
	if cfg == nil {
		return storage.NewNoopBackend(), nil
	}
	switch cfg.Storage.Backend {
	case "", "local":
		return storage.NewLocalBackend(cfg.Storage.Local.Directory)
	default:
		return storage.NewNoopBackend(), nil
	}
}
