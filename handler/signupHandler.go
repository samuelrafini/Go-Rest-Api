package handler

import (
	"PianoLessonApi/model"
	"PianoLessonApi/util"
	"database/sql"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func SignUpHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var student model.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		log.Printf("An error accured: %v", err)
		return util.NewHTTPError(err, http.StatusInternalServerError, "Server Error")
	}

	if err = util.Validate(student); err != nil {
		return util.NewHTTPError(nil, http.StatusBadRequest, err.Error())
	}

	if _, err = model.GetSingleStudentByEmail(student.Email, db); err == nil {
		return util.NewHTTPError(nil, http.StatusBadRequest, "This email already exists")
	}


	hash, err := bcrypt.GenerateFromPassword([]byte(student.Password), bcrypt.DefaultCost)
	student.Password = string(hash)

	if err = model.StoreStudent(student, db); err != nil {
		log.Printf("An error accured: %v", err)
		return util.NewHTTPError(err, http.StatusInternalServerError, "Server Error")
	}

	return nil
}
