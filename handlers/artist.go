package handlers

import (
	"fmt"
	"groupietracker/models" // Replace "yourmodule" with your actual module name

	//import needed to transform string to "North Carolina"
	"log"
	"net/http" // used to print the errors into Stdout, if necessary
	"strconv"
	"strings"
)

// ArtistDetails struct to hold data for the artist page
type ArtistDetails struct {
	Artist              models.Artist
	WikipediaURL        string
	MembersLinks        []MemberLink
	FormattedFirstAlbum string
	Title               string // Add the Title field
}

// MemberLink struct to hold data for member links
type MemberLink struct {
	Name string
	URL  string
}

// ArtistHandler displays details for a specific artist
func ArtistHandler(w http.ResponseWriter, r *http.Request, artists []models.Artist) {
	// Extract the artist ID from the URL path
	artistIDStr := strings.TrimPrefix(r.URL.Path, "/artist/")
	artistID, err := strconv.Atoi(artistIDStr)

	if err != nil || artistID < 1 || artistID > len(artists) {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	// Find the artist by ID
	artist := artists[artistID-1] // Adjust index because array start from 0 but ID's from 1

	// Create Wikipedia URL
	wikipediaURL := fmt.Sprintf("https://en.wikipedia.org/wiki/%s", strings.ReplaceAll(artist.Name, " ", "_"))

	// Create member links
	memberLinks := make([]MemberLink, len(artist.Members))
	for i, member := range artist.Members {
		memberWikiURL := fmt.Sprintf("https://en.wikipedia.org/wiki/%s", strings.ReplaceAll(member, " ", "_"))
		memberLinks[i] = MemberLink{
			Name: member,
			URL:  memberWikiURL,
		}
	}

	// Format the first album date
	formattedFirstAlbum := strings.ReplaceAll(artist.FirstAlbum, "-", ".")

	// Create the data for the template
	data := ArtistDetails{
		Artist:              artist,
		WikipediaURL:        wikipediaURL,
		MembersLinks:        memberLinks,
		FormattedFirstAlbum: formattedFirstAlbum,
		Title:               artist.Name, //here is the fix. The title was declared earlier but not used as an argument
	}

	// Execute the template
	err = tpl.ExecuteTemplate(w, "artist", data) // Render "home" template
	if err != nil {
		log.Println("Error executing template:", err)
		log.Println("Internal server error") //Output errors on console to trace back what's happening
		return
	}
}

// Handles the errors in the page
func errorHandler(w http.ResponseWriter, _ *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "<h1>Error 404:</h1><ul><li>It seems like you are lost in the woods.</li></ul>")
	}
}
