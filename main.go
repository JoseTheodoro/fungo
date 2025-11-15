package main

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("OK"))
	})

	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		a := uuid.New().String()
		message := fmt.Sprintf("Hello %s", a)
		w.Write([]byte(message))
	})
	log.Println("Server initializing 5700...")
	http.ListenAndServe(":5700", nil)

}
