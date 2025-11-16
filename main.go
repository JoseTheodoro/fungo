package main

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		a := uuid.New().String()
		message := fmt.Sprintf("Hello %s", a)
		w.Write([]byte(message))
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("OK"))
		if err != nil {
			fmt.Printf("error writing response: %v", err)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("OK"))
		if err != nil {
			fmt.Printf("error writing response: %v", err)
		}
	})

	log.Println("Server initializing 5700...")
	http.ListenAndServe(":5700", nil)

}
