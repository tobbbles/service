package cors

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.uber.org/zap"
)

func Middleware(logger *zap.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			corsConfig := cors.New(cors.Options{
				AllowedHeaders: []string{"Origin", "Accept", "Content-Type", "X-Requested-With", "Authorization"},
				AllowedMethods: []string{"POST", "PUT", "GET", "PATCH", "OPTIONS", "HEAD", "DELETE"},
				Debug:          true,
			})

			corsConfig.Handler(next).ServeHTTP(w, r)
		})
	}
}
