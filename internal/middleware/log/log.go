package log

import (
	"bytes"
	"net/http"
	"time"

	"github.com/GalichAnton/go_final_project/internal/logger"
	"go.uber.org/zap"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
	body       bytes.Buffer
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK, bytes.Buffer{}}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (lrw *loggingResponseWriter) Write(b []byte) (int, error) {
	lrw.body.Write(b)
	return lrw.ResponseWriter.Write(b)
}

func HttpLogInterceptor(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			now := time.Now()
			lrw := NewLoggingResponseWriter(w)
			next.ServeHTTP(lrw, r)

			statusCode := lrw.statusCode

			if statusCode > 399 {
				logger.Error(
					"request",
					zap.String("method", r.Method),
					zap.String("url", r.URL.String()),
					zap.Any("status", statusCode),
					zap.Duration("duration", time.Since(now)),
				)
			} else {
				logger.Info(
					"request",
					zap.String("method", r.Method),
					zap.String("url", r.URL.String()),
					zap.Any("status", statusCode),
					zap.Duration("duration", time.Since(now)),
				)
			}
		},
	)
}
