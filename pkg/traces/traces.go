package traces

import (
	"context"
	"os"
	"sync"

	gcloudtrace "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"github.com/kunitsuinc/certcounter/pkg/config"
	"github.com/kunitsuinc/certcounter/pkg/consts"
	"github.com/kunitsuinc/certcounter/pkg/errors"
	"github.com/kunitsuinc/rec.go"
	"github.com/kunitsuinc/util.go/must"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
)

type noopSpanExporter struct{}

func (noop *noopSpanExporter) ExportSpans(ctx context.Context, spans []sdktrace.ReadOnlySpan) error {
	return nil
}

func (noop *noopSpanExporter) Shutdown(ctx context.Context) error {
	return nil
}

// newExporter TODO.
func newExporter(l *rec.Logger) sdktrace.SpanExporter {
	spanExporter := config.SpanExporter()
	switch spanExporter {
	case "gcloud":
		exporter, err := gcloudtrace.New(gcloudtrace.WithProjectID(config.GoogleCloudProject()))
		if err != nil {
			l.With(rec.Error(errors.Errorf("gcloudtrace.New: %w", err))).F().Errorf("trace: %s: gcloudtrace.New: %v", spanExporter, err)
			break
		}
		return exporter
	case "stdout":
		return must.One(stdouttrace.New(stdouttrace.WithWriter(os.Stdout)))
	}

	return &noopSpanExporter{}
}

// newResource TODO.
func newResource(serviceName, version string) *resource.Resource {
	return resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(serviceName),
		semconv.ServiceVersionKey.String(version),
		semconv.TelemetrySDKLanguageGo,
	)
}

var (
	tracer      trace.Tracer    //nolint: gochecknoglobals
	tracerMutex = &sync.Mutex{} //nolint: gochecknoglobals
)

// cf. https://github.com/open-telemetry/opentelemetry-go-contrib/blob/main/instrumentation/github.com/gorilla/mux/otelmux/example/server.go
func InitTracerProvider(l *rec.Logger) (shutdown func()) {
	l.Info("trace: 🔔 start OpenTelemetry Tracer Provider")

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(newExporter(l)),
		sdktrace.WithResource(newResource(consts.AppName, config.Version())),
	)
	otel.SetTracerProvider(tracerProvider)
	tracerMutex.Lock()
	tracer = tracerProvider.Tracer(consts.TracerName)
	tracerMutex.Unlock()

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})) // TODO: TextMapPropagator について調べる

	shutdown = func() {
		defer l.Info("trace: 🔕 shutdown OpenTelemetry Tracer Provider")

		flushCtx, flushCancel := context.WithTimeout(context.Background(), config.ShutdownTimeout())
		defer flushCancel()

		if err := tracerProvider.ForceFlush(flushCtx); err != nil {
			rec.L().F().Errorf("trace: failed to flush tracer provider: %v", err)
		}

		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), config.ShutdownTimeout())
		defer shutdownCancel()

		if err := tracerProvider.Shutdown(shutdownCtx); err != nil {
			rec.L().F().Errorf("trace: failed to shutdown tracer provider: %v", err)
		}
	}

	return shutdown
}

func Start(parent context.Context, spanName string, opts ...trace.SpanStartOption) (child context.Context, span trace.Span) {
	return tracer.Start(parent, spanName, opts...)
}

func StartFunc(parent context.Context, spanName string, opts ...trace.SpanStartOption) func(spanFunction func(child context.Context) (err error)) error {
	return func(spanFunction func(context.Context) error) error {
		child, span := Start(parent, spanName, opts...)
		defer span.End()

		if err := spanFunction(child); err != nil {
			return err
		}

		return nil
	}
}
