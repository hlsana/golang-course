package main

import (
	"log"
	"net/http"

	"main/hw09/handler"
	"main/hw09/models"

	"github.com/gorilla/mux"
)

func checkAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if models.EnglishTeacher.Username != username || models.EnglishTeacher.Password != password {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func main() {
   
	r := mux.NewRouter()

    r.HandleFunc("/class", handler.GetClassInfo).Methods("GET")
    r.Handle("/student/{id}", checkAuth(http.HandlerFunc(handler.GetStudentInfo))).Methods("GET")

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatalf("Could not start server: %s\n", err.Error())
    }
}

