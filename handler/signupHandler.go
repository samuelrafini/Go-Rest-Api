package handler

import (
	"PianoLessonApi/util"
	"net/http"
)

func SignUpHandler( w http.ResponseWriter, r *http.Request) error {
	return util.NewHTTPError (nil, http.StatusCreated, "testing 123...")
}
