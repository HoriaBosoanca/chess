package main

import (
	"log"
	"os"
	"net/http"
	
	"github.com/gorilla/mux"
	"mychess/game"
)

func main() {
	// init
	router := mux.NewRouter()
	router.Use(game.EnableCORS)
	
	// get port
	port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

	// Handle endpoints
	game.HandleEndpoint(router)

	// log and start
	log.Printf("Server starting on port %s.", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}