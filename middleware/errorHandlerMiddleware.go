package middleware

import (
	"PianoLessonApi/util"
	"log"
	"net/http"
)

func ErrorHandlerMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cH, ok := h.(util.CustomHandler)
		if !ok {
			log.Printf("An error accured: %v", "CustomHandlers is not used")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err := cH.Check(w, r)
		if err == nil {
			return
		}

		clientError, ok := err.(util.CustomError)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		body, err := clientError.ResponseBody()
		if err != nil {
			log.Printf("An error accured: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		status, headers := clientError.ResponseHeaders()
		for k, v := range headers {
			w.Header().Set(k, v)
		}

		w.WriteHeader(status)
		_, e := w.Write(body)
		if e != nil {
			log.Printf("An error accured: %v", e)
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
}
