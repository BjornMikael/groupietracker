package handlers

import (
	"encoding/json"
	"fmt"
	"groupietracker/models"
	"groupietracker/utils"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// ArtistDetails struct to hold data for the artist page
type ArtistDetails struct {
	Artist              models.Artist
	WikipediaURL        string
	MembersLinks        []MemberLink
	FormattedFirstAlbum string
	Title               string
	ConcertDetails      []models.ConcertDetail
	ConcertDetailsJSON  template.JS
}

// MemberLink struct to hold data for member links
type MemberLink struct {
	Name string
}

// ConcertDetail represents a single concert event
type ConcertDetail struct {
	Date    string
	City    string
	Country string
}

// Helper function to create a list of member names
func createMemberLinks(artist models.Artist) []MemberLink {
	memberLinks := make([]MemberLink, len(artist.Members))
	for i, member := range artist.Members {
		memberLinks[i] = MemberLink{
			Name: member, // Only store the member's name
		}
	}
	return memberLinks
}

// ArtistHandler displays details for a specific artist
func ArtistHandler(w http.ResponseWriter, r *http.Request, artists []models.Artist, relations models.Relation) {
	// Extract the artist ID from the URL path
	artistIDStr := strings.TrimPrefix(r.URL.Path, "/artist/")
	artistID, err := strconv.Atoi(artistIDStr)

	if err != nil || artistID < 1 || artistID > len(artists) {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	// Find the artist by ID
	artist := artists[artistID-1]

	// Create Wikipedia URL for the band
	wikipediaURL := fmt.Sprintf("https://en.wikipedia.org/wiki/%s", strings.ReplaceAll(artist.Name, " ", "_"))

	// Format the first album date
	firstAlbumFormatted := utils.FormatDate(artist.FirstAlbum)

	// Create member links (just names, no URLs)
	memberLinks := createMemberLinks(artist)

	// Fetch concert details
	var concertDetails []models.ConcertDetail
	for _, rel := range relations.Index {
		if rel.ID == artist.ID {
			for location, dates := range rel.DatesLocations {
				// Format location into city and country
				city, country := utils.FormatLocation(location)

				// Geocode location to get latitude and longitude
				lat, lon := utils.GeocodeLocation(city + ", " + country)

				// Format dates
				for _, date := range dates {
					formattedDate := utils.FormatDate(date)
					concertDetails = append(concertDetails, models.ConcertDetail{
						Date:    formattedDate,
						City:    city,
						Country: country,
						Lat:     lat,
						Lon:     lon,
					})
				}
			}
			break
		}
	}

// Serialize concertDetails to JSON
concertDetailsJSON, err := json.Marshal(concertDetails)
if err != nil {
    log.Println("Error serializing concert details to JSON:", err)
    http.Error(w, "Internal server error", http.StatusInternalServerError)
    return
}

	// Create the data for the template
	data := ArtistDetails{
		Artist:              artist,
		WikipediaURL:        wikipediaURL,
		MembersLinks:        memberLinks,
		FormattedFirstAlbum: firstAlbumFormatted,
		Title:               artist.Name,
		ConcertDetails:      concertDetails,
		ConcertDetailsJSON:  template.JS(concertDetailsJSON), // Add the JSON string to the template data
	}

	// Execute the template
	err = tpl.ExecuteTemplate(w, "artist", data)
	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
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
