package handler

import (
	"PianoLessonApi/util"
	"net/http"
)

func SignUpHandler( w http.ResponseWriter, r *http.Request) error {
	return util.NewHTTPError (nil, 405, "testing 123...")
	//return nil
}
