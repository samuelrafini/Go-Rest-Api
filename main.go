package main

import (
	"PianoLessonApi/handler"
	"PianoLessonApi/middleware"
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
	fmt.Println(serverPort)

	r := mux.NewRouter()
	//r.Use(middleware.AuthMiddleware)
	r.Handle("/signup", middleware.ErrorHandlerWrapper(handler.SignUpHandler)).Methods("POST")

	if err := http.ListenAndServe(serverPort, r); err != nil {
		panic(err)
	}
}

