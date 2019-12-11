package middleware

import "net/http"

type ErrorHandlerWrapper func(w http.ResponseWriter, r *http.Request) error

func (hw ErrorHandlerWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := hw(w, r); if err == nil {
		return
	}
	w.WriteHeader(305) //testing
}
