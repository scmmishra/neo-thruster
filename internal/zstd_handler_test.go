package internal

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/klauspost/compress/zstd"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestZstdHandler(t *testing.T) {
	handler := ZstdHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	}))

	t.Run("compresses when zstd is accepted", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/", nil)
		r.Header.Set("Accept-Encoding", "zstd")

		handler.ServeHTTP(w, r)

		assert.Equal(t, "zstd", w.Header().Get("Content-Encoding"))

		decoder, err := zstd.NewReader(w.Body)
		require.NoError(t, err)
		defer decoder.Close()

		content, err := io.ReadAll(decoder)
		require.NoError(t, err)
		assert.Equal(t, "hello world", string(content))
	})

	t.Run("skips compression when zstd not accepted", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/", nil)

		handler.ServeHTTP(w, r)

		assert.Empty(t, w.Header().Get("Content-Encoding"))
		assert.Equal(t, "hello world", w.Body.String())
	})

	t.Run("skips compression when already encoded", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/", nil)
		r.Header.Set("Accept-Encoding", "zstd")
		w.Header().Set("Content-Encoding", "gzip")

		handler.ServeHTTP(w, r)

		assert.Equal(t, "gzip", w.Header().Get("Content-Encoding"))
	})
}
