package main

import (
	"log"
	"os"
	"net/http"
	
	"github.com/gorilla/mux"

	"chess/chess"
)

func main() {
	// init
	router := mux.NewRouter()
	router.Use(chess.EnableCORS)
	
	// get port
	port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

	// Handle endpoints
	chess.HandleEndpoint(router)

	// log and start
	log.Printf("Server starting on port %s.", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}