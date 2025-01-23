package internal

import (
	"net/http"
	"strings"

	"github.com/klauspost/compress/zstd"
)

type zstdResponseWriter struct {
	http.ResponseWriter
	writer *zstd.Encoder
}

func (w *zstdResponseWriter) Write(b []byte) (int, error) {
	return w.writer.Write(b)
}

// ZstdHandler wraps an http.Handler to support zstd compression
func ZstdHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "zstd") {
			h.ServeHTTP(w, r)
			return
		}

		// Don't compress if already encoded
		if w.Header().Get("Content-Encoding") != "" {
			h.ServeHTTP(w, r)
			return
		}

		encoder, err := zstd.NewWriter(w)
		if err != nil {
			h.ServeHTTP(w, r)
			return
		}
		defer encoder.Close()

		w.Header().Set("Content-Encoding", "zstd")
		w.Header().Del("Content-Length")

		zw := &zstdResponseWriter{
			ResponseWriter: w,
			writer:         encoder,
		}

		h.ServeHTTP(zw, r)
	})
}
