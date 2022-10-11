package config

import (
	"flag"
	"sync"
	"time"

	"github.com/kunitsuinc/certcounter/pkg/errors"
	"github.com/kunitsuinc/rec.go"
	"github.com/kunitsuinc/util.go/env"
)

var ErrFlagOrEnvIsNotEnough = errors.New("config: flag or environment variable is not enough")

// nolint: revive,stylecheck
const (
	DEBUG                = "DEBUG"
	APP_ENV              = "APP_ENV"
	ADDR                 = "ADDR"
	PORT                 = "PORT"
	GRPC_ENDPOINT        = "GRPC_ENDPOINT"
	SHUTDOWN_TIMEOUT     = "SHUTDOWN_TIMEOUT"
	AWS_PROFILE          = "AWS_PROFILE"
	GOOGLE_CLOUD_PROJECT = "GOOGLE_CLOUD_PROJECT"
	SPAN_EXPORTER        = "SPAN_EXPORTER"
)

type Config struct {
	Debug              bool
	AppEnv             string
	Addr               string
	Port               int
	GRPCEndpoint       string
	ShutdownTimeout    time.Duration
	AWSProfile         string
	GoogleCloudProject string
	SpanExporter       string
}

// nolint: gochecknoglobals
var (
	cfg   = &Config{}
	cfgMu sync.Mutex
)

func Load(l *rec.Logger) {
	cfgMu.Lock()
	defer cfgMu.Unlock()

	var timeout int64

	// version
	flag.BoolVar(&subcommandVersion, "v", false, "Display version info")
	flag.BoolVar(&subcommandVersion, "version", subcommandVersion || false, "Display version info")
	// config
	flag.BoolVar(&cfg.Debug, "debug", env.BoolOrDefault(DEBUG, false), "Debug")
	flag.StringVar(&cfg.AppEnv, "appenv", env.StringOrDefault(APP_ENV, ""), "TODO")
	flag.StringVar(&cfg.Addr, "addr", env.StringOrDefault(ADDR, "0.0.0.0"), "TODO")
	flag.IntVar(&cfg.Port, "port", env.IntOrDefault(PORT, 8080), "TODO")
	flag.StringVar(&cfg.GRPCEndpoint, "grpc-endpoint", env.StringOrDefault(ADDR, "0.0.0.0:9090"), "TODO")
	flag.Int64Var(&timeout, "shutdown-timeout", env.Int64OrDefault(SHUTDOWN_TIMEOUT, 10), "TODO")
	flag.StringVar(&cfg.AWSProfile, "aws-profile", env.StringOrDefault(AWS_PROFILE, ""), "TODO")
	flag.StringVar(&cfg.GoogleCloudProject, "gcp-project", env.StringOrDefault(GOOGLE_CLOUD_PROJECT, ""), "TODO")
	flag.StringVar(&cfg.SpanExporter, "span-exporter", env.StringOrDefault(SPAN_EXPORTER, "gcloud"), "TODO")
	// parse
	flag.Parse()

	cfg.ShutdownTimeout = time.Duration(timeout) * time.Second

	l.Info("config: loaded", rec.Object("config", cfg))
}

// nolint: cyclop
func Check() error {
	switch {
	case cfg.Addr == "":
		return errors.Errorf("%s: %w", ADDR, ErrFlagOrEnvIsNotEnough)
	case cfg.Port == 0:
		return errors.Errorf("%s: %w", PORT, ErrFlagOrEnvIsNotEnough)
	case cfg.GRPCEndpoint == "":
		return errors.Errorf("%s: %w", GRPC_ENDPOINT, ErrFlagOrEnvIsNotEnough)
	}

	switch {
	case cfg.AWSProfile != "":
		break
	case cfg.GoogleCloudProject != "":
		break
	default:
		return errors.Errorf("%s || %s: %w", AWS_PROFILE, GOOGLE_CLOUD_PROJECT, ErrFlagOrEnvIsNotEnough)
	}

	return nil
}

func Debug() bool                    { return cfg.Debug }
func AppEnv() string                 { return cfg.AppEnv }
func Addr() string                   { return cfg.Addr }
func Port() int                      { return cfg.Port }
func GRPCEndpoint() string           { return cfg.GRPCEndpoint }
func ShutdownTimeout() time.Duration { return cfg.ShutdownTimeout }
func AWSProfile() string             { return cfg.AWSProfile }
func GoogleCloudProject() string     { return cfg.GoogleCloudProject }
func SpanExporter() string           { return cfg.SpanExporter }
