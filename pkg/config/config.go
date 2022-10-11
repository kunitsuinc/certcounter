package config

import (
	"flag"
	"os"
	"sync"
	"time"

	"github.com/kunitsuinc/certcounter/pkg/errors"
	"github.com/kunitsuinc/rec.go"
	"github.com/kunitsuinc/util.go/env"
	"github.com/kunitsuinc/util.go/pointer"
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
	cfgBackup = &Config{}
	cfg       = &Config{}
	cfgMu     sync.Mutex
)

func Load(l *rec.Logger) (rollback func(), err error) {
	cfgMu.Lock()
	defer cfgMu.Unlock()
	cfgBackup = pointer.Pointer(*cfg)

	rollback = func() {
		cfgMu.Lock()
		defer cfgMu.Unlock()
		cfg = pointer.Pointer(*cfgBackup)
		l.Info("config: loaded", rec.Object("config", cfg))
	}

	var timeout int64

	fs := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	// version
	fs.BoolVar(&subcommandVersion, "v", false, "Display version info")
	fs.BoolVar(&subcommandVersion, "version", subcommandVersion || false, "Display version info")
	// config
	fs.BoolVar(&cfg.Debug, "debug", env.BoolOrDefault(DEBUG, false), "Debug")
	fs.StringVar(&cfg.AppEnv, "appenv", env.StringOrDefault(APP_ENV, ""), "TODO")
	fs.StringVar(&cfg.Addr, "addr", env.StringOrDefault(ADDR, "0.0.0.0"), "TODO")
	fs.IntVar(&cfg.Port, "port", env.IntOrDefault(PORT, 8080), "TODO")
	fs.StringVar(&cfg.GRPCEndpoint, "grpc-endpoint", env.StringOrDefault(ADDR, "0.0.0.0:9090"), "TODO")
	fs.Int64Var(&timeout, "shutdown-timeout", env.Int64OrDefault(SHUTDOWN_TIMEOUT, 10), "TODO")
	fs.StringVar(&cfg.AWSProfile, "aws-profile", env.StringOrDefault(AWS_PROFILE, ""), "TODO")
	fs.StringVar(&cfg.GoogleCloudProject, "gcp-project", env.StringOrDefault(GOOGLE_CLOUD_PROJECT, ""), "TODO")
	fs.StringVar(&cfg.SpanExporter, "span-exporter", env.StringOrDefault(SPAN_EXPORTER, "gcloud"), "TODO")
	// parse
	if err := fs.Parse(os.Args[1:]); err != nil {
		return rollback, errors.Errorf("(*flag.FlagSet).Parse: %w", err)
	}

	cfg.ShutdownTimeout = time.Duration(timeout) * time.Second

	l.Info("config: loaded", rec.Object("config", cfg))

	switch {
	case cfg.Addr == "":
		return rollback, errors.Errorf("%s: %w", ADDR, ErrFlagOrEnvIsNotEnough)
	case cfg.Port == 0:
		return rollback, errors.Errorf("%s: %w", PORT, ErrFlagOrEnvIsNotEnough)
	case cfg.GRPCEndpoint == "":
		return rollback, errors.Errorf("%s: %w", GRPC_ENDPOINT, ErrFlagOrEnvIsNotEnough)
	}

	switch {
	case cfg.AWSProfile != "":
		break
	case cfg.GoogleCloudProject != "":
		break
	default:
		return rollback, errors.Errorf("%s || %s: %w", AWS_PROFILE, GOOGLE_CLOUD_PROJECT, ErrFlagOrEnvIsNotEnough)
	}

	return rollback, nil
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
