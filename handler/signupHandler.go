package handler

import (
	"PianoLessonApi/util"
	"database/sql"
	"net/http"
)

func SignUpHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	return util.NewHTTPError (nil, http.StatusCreated, "testing 123...")
}
