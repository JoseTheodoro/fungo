package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func main() {

	http.HandleFunc("/api/deploy", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		message := fmt.Sprintf("Deployed %s", uuid.New().String())
		_, err := w.Write([]byte(message))
		if err != nil {
			fmt.Printf("error writing response: %v", err)
		}
	})

	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		a := uuid.New().String()
		message := fmt.Sprintf("Hello %s", a)
		_, err := w.Write([]byte(message))
		if err != nil {
			fmt.Printf("error writing response: %v", err)
		}
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		a := uuid.New().String()
		w.Write([]byte(a))
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		message := fmt.Sprintf("OK %s", uuid.New().String())
		_, err := w.Write([]byte(message))
		if err != nil {
			fmt.Printf("error writing response: %v", err)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		message := fmt.Sprintf("OK %s", uuid.New().String())
		_, err := w.Write([]byte(message))
		if err != nil {
			fmt.Printf("error writing response: %v", err)
		}
	})

	log.Println("Server initializing 5700...")
	http.ListenAndServe(":5700", nil)

}
