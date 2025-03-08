package main

import (
	"fmt"
	"groupietracker/handlers"
	"groupietracker/models"
	"groupietracker/utils"
	"log"
	"net/http"
)

var artists []models.Artist
var relations models.Relation

func main() {
	// Fetch artists data
	err := utils.GetArtists(&artists)
	if err != nil {
		log.Println("Error fetching artists:", err)
	}

	// Fetch relations data
	err = utils.GetRelations(&relations)
	if err != nil {
		log.Println("Error fetching relations:", err)
	}

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Home route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HomeHandler(w, r, artists)
	})

	// Artist route
	http.HandleFunc("/artist/", func(w http.ResponseWriter, r *http.Request) {
		handlers.ArtistHandler(w, r, artists, relations) // Pass relations here
	})

	// Start the server
	fmt.Println("Server listening on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
