package handler

import (
	"net/http"

	"github.com/lukebull/go-zero-extern/core/logx"
	"github.com/lukebull/go-zero-extern/core/syncx"
	"github.com/lukebull/go-zero-extern/rest/internal"
)

// MaxConns returns a middleware that limit the concurrent connections.
func MaxConns(n int) func(http.Handler) http.Handler {
	if n <= 0 {
		return func(next http.Handler) http.Handler {
			return next
		}
	}

	return func(next http.Handler) http.Handler {
		latch := syncx.NewLimit(n)

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if latch.TryBorrow() {
				defer func() {
					if err := latch.Return(); err != nil {
						logx.Error(err)
					}
				}()

				next.ServeHTTP(w, r)
			} else {
				internal.Errorf(r, "concurrent connections over %d, rejected with code %d",
					n, http.StatusServiceUnavailable)
				w.WriteHeader(http.StatusServiceUnavailable)
			}
		})
	}
}
