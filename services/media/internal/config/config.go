package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/posflag"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/knadh/koanf/v2"
	"github.com/spf13/pflag"
)

type Config struct {
	Log        LogConfig        `koanf:"log"`
	Server     ServerConfig     `koanf:"server"`
	Storage    StorageConfig    `koanf:"storage"`
	Processing ProcessingConfig `koanf:"processing"`
	Services   ServicesConfig   `koanf:"services"`
	Debug      DebugConfig      `koanf:"debug"`
}

type LogConfig struct {
	Level string `koanf:"level"`
}

type ServerConfig struct {
	Port string `koanf:"port"`
}

type StorageConfig struct {
	Backend string             `koanf:"backend"`
	Local   LocalStorageConfig `koanf:"local"`
	S3      S3StorageConfig    `koanf:"s3"`
}

type LocalStorageConfig struct {
	Directory string `koanf:"directory"`
}

type S3StorageConfig struct {
	Bucket    string `koanf:"bucket"`
	Region    string `koanf:"region"`
	Endpoint  string `koanf:"endpoint"`
	AccessKey string `koanf:"access_key"`
	SecretKey string `koanf:"secret_key"`
}

type ProcessingConfig struct {
	Cropping    bool            `koanf:"cropping"`
	Compression bool            `koanf:"compression"`
	Variants    []VariantConfig `koanf:"variants"`
}

type VariantConfig struct {
	Name   string `koanf:"name"`
	Width  int    `koanf:"width"`
	Height int    `koanf:"height"`
}

type ServicesConfig struct {
	DictionaryURL string `koanf:"dictionary_url"`
}

type DebugConfig struct {
	Routes bool `koanf:"routes"`
}

func New() *Config {
	return &Config{
		Log:    LogConfig{Level: "info"},
		Server: ServerConfig{Port: ":8086"},
		Storage: StorageConfig{
			Backend: "local",
			Local: LocalStorageConfig{
				Directory: "./data/media",
			},
		},
		Processing: ProcessingConfig{
			Cropping:    false,
			Compression: false,
			Variants: []VariantConfig{
				{Name: "thumb", Width: 150, Height: 150},
				{Name: "preview", Width: 800, Height: 600},
				{Name: "web", Width: 1280, Height: 720},
			},
		},
		Services: ServicesConfig{
			DictionaryURL: "http://localhost:8082",
		},
		Debug: DebugConfig{Routes: true},
	}
}

func LoadConfig(path, envPrefix string, args []string) (*Config, error) {
	k := koanf.New(".")
	cfg := New()

	fs := pflag.NewFlagSet(args[0], pflag.ExitOnError)
	fs.String("server.port", cfg.Server.Port, "Server listen address")
	fs.String("log.level", cfg.Log.Level, "Log level")
	fs.String("storage.backend", cfg.Storage.Backend, "Storage backend (local|s3)")
	fs.String("storage.local.directory", cfg.Storage.Local.Directory, "Local storage directory")
	fs.String("services.dictionary_url", cfg.Services.DictionaryURL, "Dictionary service base URL")
	fs.Bool("debug.routes", cfg.Debug.Routes, "Expose /debug/routes endpoint")
	fs.Bool("processing.cropping", cfg.Processing.Cropping, "Enable automatic cropping")
	fs.Bool("processing.compression", cfg.Processing.Compression, "Enable compression")
	fs.Parse(args[1:])

	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read config file: %w", err)
	}
	expanded := []byte(os.ExpandEnv(string(raw)))

	if err := k.Load(rawbytes.Provider(expanded), yaml.Parser()); err != nil {
		return nil, fmt.Errorf("cannot parse yaml: %w", err)
	}

	if err := k.Load(env.Provider(envPrefix, ".", func(s string) string {
		return strings.Replace(strings.ToLower(strings.TrimPrefix(s, envPrefix)), "_", ".", -1)
	}), nil); err != nil {
		return nil, fmt.Errorf("cannot load env vars: %w", err)
	}

	if err := k.Load(posflag.Provider(fs, ".", k), nil); err != nil {
		return nil, fmt.Errorf("cannot load flags: %w", err)
	}

	if err := k.Unmarshal("", cfg); err != nil {
		return nil, fmt.Errorf("cannot unmarshal config: %w", err)
	}

	return cfg, nil
}
