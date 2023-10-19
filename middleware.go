package bot

import (
	"net/http"

	"github.com/google/uuid"
)

func RequestID(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		ctx := SetRequestID(r.Context(), id)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

type loggingResponseWriter struct {
	w      http.ResponseWriter
	status int
}

func (w *loggingResponseWriter) Header() http.Header {
	return w.w.Header()
}

func (w *loggingResponseWriter) Write(b []byte) (int, error) {
	return w.w.Write(b)
}

func (w *loggingResponseWriter) WriteHeader(staturCode int) {
	w.status = staturCode
	w.w.WriteHeader(staturCode)
}

func Logging(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := SetRequestRecord(r.Context(), r)
		defer func() {
			Info(ctx, "finish process")
		}()
		r = r.WithContext(ctx)

		lw := &loggingResponseWriter{
			w:      w,
			status: http.StatusOK,
		}
		h.ServeHTTP(lw, r)
		ctx = SetResponseRecord(r.Context(), lw.status)
	}

	return http.HandlerFunc(fn)
}
