package config

import (
	"flag"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/kunitsuinc/rec.go"
	"github.com/kunitsuinc/util.go/flagenv"
	"github.com/kunitsuinc/util.go/pointer"

	v1 "github.com/kunitsuinc/certcounter/generated/go/certcounter/v1"
	"github.com/kunitsuinc/certcounter/pkg/errors"
)

var ErrFlagOrEnvIsNotEnough = errors.New("config: flag or environment variable is not enough")

//nolint:revive,stylecheck
const (
	VERSION              = "VERSION"
	DEBUG                = "DEBUG"
	APP_ENV              = "APP_ENV"
	ADDR                 = "ADDR"
	PORT                 = "PORT"
	SET_REAL_IP_FROM     = "SET_REAL_IP_FROM"
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
	SetRealIPFrom      []*net.IPNet
	ShutdownTimeout    time.Duration
	AWSProfile         string
	GoogleCloudProject string
	SpanExporter       string
}

//nolint:gochecknoglobals
var (
	cfgBackup = &Config{}
	cfg       = &Config{}
	cfgMu     sync.Mutex
)

// Load
//
//nolint:funlen
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

	var setRealIPFrom string

	fs := flagenv.NewFlagEnvSet(os.Args[0], flag.ContinueOnError)
	// version
	fs.BoolVar(&subcommandVersion, "v", VERSION, false, "Display version info")
	fs.BoolVar(&subcommandVersion, "version", VERSION, subcommandVersion || false, "Display version info")
	// config
	fs.BoolVar(&cfg.Debug, "debug", DEBUG, false, "Debug")
	fs.StringVar(&cfg.AppEnv, "appenv", APP_ENV, "", "TODO")
	fs.StringVar(&cfg.Addr, "addr", ADDR, "0.0.0.0", "TODO")
	fs.IntVar(&cfg.Port, "port", PORT, 8080, "TODO")
	fs.StringVar(&setRealIPFrom, "set-real-ip-from", SET_REAL_IP_FROM, "10.0.0.0/8,172.16.0.0/12,192.168.0.0/16", "TODO")
	fs.SecondVar(&cfg.ShutdownTimeout, "shutdown-timeout", SHUTDOWN_TIMEOUT, 10*time.Second, "TODO")
	fs.StringVar(&cfg.AWSProfile, "aws-profile", AWS_PROFILE, "", "TODO")
	fs.StringVar(&cfg.GoogleCloudProject, "gcp-project", GOOGLE_CLOUD_PROJECT, "", "TODO")
	fs.StringVar(&cfg.SpanExporter, "span-exporter", SPAN_EXPORTER, v1.SpanExporter_stdout.String(), "TODO")
	// parse
	if err := fs.Parse(os.Args[1:]); err != nil {
		return rollback, errors.Errorf("(*flag.FlagSet).Parse: %w", err)
	}

	// set_real_ip_from
	for _, v := range strings.Split(setRealIPFrom, ",") {
		ip, ipNet, err := net.ParseCIDR(strings.TrimSpace(v))
		if err != nil {
			err = errors.Errorf("net.ParseCIDR: %w")
			l.With(rec.Error(err)).E().Warning(err)
			continue
		}
		ipNet.IP = ip
		cfg.SetRealIPFrom = append(cfg.SetRealIPFrom, ipNet)
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
func SetRealIPFrom() []*net.IPNet    { return cfg.SetRealIPFrom }
func ShutdownTimeout() time.Duration { return cfg.ShutdownTimeout }
func AWSProfile() string             { return cfg.AWSProfile }
func GoogleCloudProject() string     { return cfg.GoogleCloudProject }
func SpanExporter() string           { return cfg.SpanExporter }
