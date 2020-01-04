package util

import (
	"database/sql"
	"log"
	"net/http"
)

type CustomHandler interface {
 Check(http.ResponseWriter, *http.Request) error
}

type Handler struct {
	DB     *sql.DB
	Handle func(db *sql.DB, w http.ResponseWriter, r *http.Request) error
}

func (h Handler) Check(w http.ResponseWriter, r *http.Request) error {
	err := h.Handle(h.DB, w, r)
	if err != nil {
		return err
	}
	return nil
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.Handle(h.DB, w, r)
	if err != nil {
		log.Printf("An error accured: %v", err)
	}
}
