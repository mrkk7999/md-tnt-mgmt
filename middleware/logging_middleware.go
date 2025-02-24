package middleware

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// LoggingMiddleware logs the incoming request details and completion time
func LoggingMiddleware(next http.Handler, log *logrus.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		log.Info("-----------------------------------------------------------")
		log.Infof("Incoming request: Method=%s URL=%s Timestamp=%s", r.Method, r.URL.Path, startTime.Format(time.RFC3339))

		// Call next handler
		next.ServeHTTP(w, r)
		log.Infof("Request Completed: URL=%s ", r.URL.Path)
		log.Info("-----------------------------------------------------------")

	})
}
