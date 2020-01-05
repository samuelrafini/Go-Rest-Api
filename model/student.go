package model

import (
	"database/sql"
	"fmt"
	"log"
)

type Student struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

func StoreStudent(student Student, db *sql.DB) error {
	var stmt = "insert into students (first_name, last_name, email, password) values($1, $2, $3, $4) returning email"
	err := db.QueryRow(stmt, student.FirstName, student.LastName, student.Email, student.Password).Scan(&student.Email)
	log.Print(err)
	if err != nil {
		return err
	}
	return nil
}

func GetSingleStudentByEmail(email string, db *sql.DB) (*Student, error) {
	var stmt = "select first_name, last_name, email from students where email=$1"
	var student Student
	err := db.QueryRow(stmt, email).Scan(&student.FirstName, &student.LastName ,&student.Email)
	if err != nil {
		fmt.Println(err)
		return nil , err
	}
	return &student, nil
}
