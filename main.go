package main

import (
	"PianoLessonApi/handler"
	"PianoLessonApi/middleware"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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
	port := os.Getenv("PORT")
	fmt.Println(port)

	r := mux.NewRouter()
	r.Handle("/signup", middleware.ErrorHandlerWrapper(handler.SignUpHandler)).Methods("POST")

	if err := http.ListenAndServe(port, r); err != nil {
		panic(err)
	}
}

