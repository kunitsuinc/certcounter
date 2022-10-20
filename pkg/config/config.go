package config

import (
	"flag"
	"os"
	"sync"
	"time"

	"github.com/kunitsuinc/certcounter/pkg/errors"
	"github.com/kunitsuinc/rec.go"
	"github.com/kunitsuinc/util.go/flagenv"
	"github.com/kunitsuinc/util.go/pointer"
)

var ErrFlagOrEnvIsNotEnough = errors.New("config: flag or environment variable is not enough")

// nolint: revive,stylecheck
const (
	VERSION              = "VERSION"
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

	fs := flagenv.NewFlagEnvSet(os.Args[0], flag.ContinueOnError)
	// version
	fs.BoolVar(&subcommandVersion, "v", VERSION, false, "Display version info")
	fs.BoolVar(&subcommandVersion, "version", VERSION, subcommandVersion || false, "Display version info")
	// config
	fs.BoolVar(&cfg.Debug, "debug", DEBUG, false, "Debug")
	fs.StringVar(&cfg.AppEnv, "appenv", APP_ENV, "", "TODO")
	fs.StringVar(&cfg.Addr, "addr", ADDR, "0.0.0.0", "TODO")
	fs.IntVar(&cfg.Port, "port", PORT, 8080, "TODO")
	fs.SecondVar(&cfg.ShutdownTimeout, "shutdown-timeout", SHUTDOWN_TIMEOUT, 10*time.Second, "TODO")
	fs.StringVar(&cfg.AWSProfile, "aws-profile", AWS_PROFILE, "", "TODO")
	fs.StringVar(&cfg.GoogleCloudProject, "gcp-project", GOOGLE_CLOUD_PROJECT, "", "TODO")
	fs.StringVar(&cfg.SpanExporter, "span-exporter", SPAN_EXPORTER, "gcloud", "TODO")
	// parse
	if err := fs.Parse(os.Args[1:]); err != nil {
		return rollback, errors.Errorf("(*flag.FlagSet).Parse: %w", err)
	}

	l.Info("config: loaded", rec.Object("config", cfg))

	// AND
	switch {
	case cfg.Addr == "":
		return rollback, errors.Errorf("%s: %w", ADDR, ErrFlagOrEnvIsNotEnough)
	case cfg.Port == 0:
		return rollback, errors.Errorf("%s: %w", PORT, ErrFlagOrEnvIsNotEnough)
	}

	// OR
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
func ShutdownTimeout() time.Duration { return cfg.ShutdownTimeout }
func AWSProfile() string             { return cfg.AWSProfile }
func GoogleCloudProject() string     { return cfg.GoogleCloudProject }
func SpanExporter() string           { return cfg.SpanExporter }
