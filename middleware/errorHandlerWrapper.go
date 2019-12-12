package middleware

import (
	"log"
	"net/http"
	)

type ClientError interface {
	Error() string
	ResponseBody() ([]byte, error)
	ResponseHeaders() (int, map[string]string)
}

type ErrorHandlerWrapper func(w http.ResponseWriter, r *http.Request) error

func (hw ErrorHandlerWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := hw(w, r); if err == nil {
		return
	}

	log.Printf("An error accured: %v", err)

	clientError, ok := err.(ClientError)
	if !ok {
		w.WriteHeader(500)
		return
	}

	body, err := clientError.ResponseBody()
	if err != nil {
		log.Printf("An error accured: %v", err)
		w.WriteHeader(500)
		return
	}

	status, headers := clientError.ResponseHeaders()
	for k, v := range headers {
		w.Header().Set(k, v)
	}

	w.WriteHeader(status)
	_, e := w.Write(body); if e != nil {
		log.Printf("An error accured: %v", e)
	}
}
