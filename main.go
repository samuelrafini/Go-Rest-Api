package main

import (
	"PianoLessonApi/handler"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)



//func handleSignup(w http.ResponseWriter, r *http.Request) error {
//	json.NewDecoder(r.Body)
//}

type handleWrapper func(w http.ResponseWriter, r *http.Request) error

func (hw handleWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := hw(w, r); if err == nil {
		return
	}
	w.WriteHeader(305) //testing
}

//func errorMiddleware(h http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
//		body, _ := ioutil.ReadAll(r.Body)
//		fmt.Fprint(w, string(body))
//		fmt.Println("test it is working")
//		//err := h
//		//fmt.Println()
//		//if err != nil {
//		//	fmt.Fprint(w, 400)
//		//	w.WriteHeader(400)
//		//	//w.Header().Set()
//		//	return
//		//}
//		//w.WriteHeader(err.)
//	})
//}


func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	port := os.Getenv("PORT")
	fmt.Println(port)

	r := mux.NewRouter()
	//r.Handle("/signup", errorMiddleware(handler.HandleSignup()))
	r.Handle("/signup", handleWrapper(handler.SignUpHandler))
	//r.Use(errorMiddleware)
	//r.HandleFunc("/signup", errorMiddleware()).Methods("POST")

	if err := http.ListenAndServe(port, r); err != nil {
		panic(err)
	}
}

