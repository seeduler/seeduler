package middlewares

import (
	"net/http"
	"github.com/seeduler/seeduler/services"
)

func HallCheckMiddleware(hallService *services.HallService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Allow upload-data endpoint to bypass this check
			if r.URL.Path == "/halls/upload-data" {
				next.ServeHTTP(w, r)
				return
			}

			// Check if halls exist
			halls, err := hallService.GetAllHalls()
			if err != nil || len(halls) == 0 {
				http.Error(w, "No halls found. Please upload data first.", http.StatusPreconditionFailed)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
} 