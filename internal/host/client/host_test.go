package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHTTPDo(t *testing.T) {
	t.Run("request success", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte("Hello, client!"))
		}))
		defer srv.Close()

		statusCode, err := (host{}).HTTPDo(context.Background(), "GET", srv.URL, nil, nil)
		require.NoError(t, err)
		require.Equal(t, http.StatusAccepted, statusCode)
	})

	t.Run("request fails", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			panic("test")
		}))
		defer srv.Close()

		statusCode, err := (host{}).HTTPDo(context.Background(), "GET", srv.URL, nil, nil)
		require.Error(t, err)
		require.Zero(t, statusCode)
	})
}
