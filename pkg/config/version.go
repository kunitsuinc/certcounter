package config

import (
	"runtime/debug"
)

//nolint:deadcode,gochecknoglobals,unused,varcheck
var (
	subcommandVersion bool
	version           string
	revision          string
	branch            string
	timestamp         string
)

//nolint:gochecknoglobals
var goVersion = func() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "go"
	}
	return info.GoVersion
}()

func SubcommandVersion() bool { return subcommandVersion }
func GoVersion() string       { return goVersion }
func Version() string         { return version }
func Revision() string        { return revision }
func Branch() string          { return branch }
func Timestamp() string       { return timestamp }
