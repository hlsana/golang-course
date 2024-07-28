package main

import (
	"log"
	"main/hw10/router"
	"net/http"
)

func main() {
	r := router.InitializeRouter()

	log.Println("Server started on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
