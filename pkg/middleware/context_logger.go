package middleware

import (
	"net/http"

	"github.com/kunitsuinc/rec.go"
	"github.com/kunitsuinc/util.go/net/httpz"
)

func ContextLoggerRequestMiddleware(original *rec.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(rw, r.WithContext(rec.ContextWithLogger(r.Context(), original.With(
				rec.String("remoteAddr", r.RemoteAddr),
				rec.String("xRealIP", r.Header.Get(httpz.HeaderXRealIP)),
				rec.String("method", r.Method),
				rec.String("url", r.URL.String()),
				rec.String("host", r.Host),
				rec.String("path", r.URL.Path),
				rec.String("query", r.URL.RawQuery),
				rec.String("proto", r.Proto),
				rec.Int64("requestContentLength", r.ContentLength),
				// rec.String("referer", r.Referer()),
				rec.String("userAgent", r.UserAgent()),
			))))
		})
	}
}
