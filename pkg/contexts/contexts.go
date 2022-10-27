package contexts

import (
	"context"
	"os"

	"github.com/kunitsuinc/certcounter/pkg/errors"
)

var ErrValueNotSet = errors.New("contexts: context value not set")

type key int

//nolint:gochecknoglobals
const (
	_ key = iota
	shutdownChanKey
)

func WithShutdownChan(parent context.Context, ch chan os.Signal) context.Context {
	return context.WithValue(parent, shutdownChanKey, ch)
}

func ShutdownChan(ctx context.Context) chan os.Signal {
	ch, ok := ctx.Value(shutdownChanKey).(chan os.Signal)
	if ok && ch != nil {
		return ch
	}

	panic(errors.Errorf("value=ShutdownChan: %w", ErrValueNotSet))
}
