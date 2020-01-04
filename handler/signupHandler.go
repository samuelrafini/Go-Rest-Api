package handler

import (
	"PianoLessonApi/model"
	"PianoLessonApi/util"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func SignUpHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var student model.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		log.Printf("An error accured: %v", err)
		return util.NewHTTPError(err, http.StatusInternalServerError, "")
	}
	if err = util.Validate(student); err != nil {
		return util.NewHTTPError(nil, http.StatusBadRequest, err.Error())
	}
	return nil
}
