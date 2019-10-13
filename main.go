package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bramalho/jira-tasks/controller"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", controller.HomeHandler).Methods("GET")

	router.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	port, exists := os.LookupEnv("PORT")
	if exists == false {
		port = "8080"
	}

	srv := &http.Server{
		Handler:      router,
		Addr:         "localhost:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Application is running on localhost:" + port)

	log.Fatal(srv.ListenAndServe())
}
