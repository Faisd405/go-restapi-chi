package middleware

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func makeLogEntry(r *http.Request) *log.Entry {
	return log.WithFields(log.Fields{
		"Time":   time.Now().Format(time.RFC3339),
		"Method": r.Method,
		"URI":    r.RequestURI,
		"IP":     r.RemoteAddr,
	})
}

func MiddlewareLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		makeLogEntry(r).Info("request received")
		next.ServeHTTP(w, r)
	})
}

func ErrorHandler(err error, r *http.Request) {
	makeLogEntry(r).Error(err)
}
