package main

import (
	"fmt"
	"groupietracker/handlers" // Import the handlers package
	"groupietracker/models"   // Replace "yourmodule" with your actual module name
	"groupietracker/utils"    // Replace "yourmodule" with your actual module name
	"log"
	"net/http"
)

// Define a global variable to store the artists data
var artists []models.Artist

func main() {
	var locations models.Locations
	var dates models.Dates
	var relations models.Relation

	err := utils.GetArtists(&artists)
	if err != nil {
		log.Println("Error fetching artists:", err)
	}

	err = utils.GetLocations(&locations)
	if err != nil {
		log.Println("Error fetching locations:", err)
	}

	err = utils.GetDates(&dates)
	if err != nil {
		log.Println("Error fetching dates:", err)
	}

	err = utils.GetRelations(&relations)
	if err != nil {
		log.Println("Error fetching relations:", err)
	}

	// Serve static files (CSS, JavaScript, images)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// ---- WEB SERVER SETUP ----
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //pass artists data in the homeHandler function
		handlers.HomeHandler(w, r, artists)
	})

	// Register the artist handler for the "/artist/{id}" route
	http.HandleFunc("/artist/", func(w http.ResponseWriter, r *http.Request) {
		handlers.ArtistHandler(w, r, artists)
	})

	fmt.Println("Server listening on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
