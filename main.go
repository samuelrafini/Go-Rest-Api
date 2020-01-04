package main

import (
	"PianoLessonApi/handler"
	"PianoLessonApi/middleware"
	"PianoLessonApi/util"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	serverPort := os.Getenv("SERVER_PORT")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := util.ConnectDB(psqlInfo)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	log.Println(serverPort)

	r := mux.NewRouter()
	r.Handle("/signup", util.Handler{db, handler.SignUpHandler}).Methods("POST")
	r.Use(middleware.ErrorHandlerMiddleware)

	if err := http.ListenAndServe(serverPort, r); err != nil {
		log.Println(err)
	}
}
