package utils

import (
	"encoding/json"
	"fmt"
	"groupietracker/models" // Replace "yourmodule" with your actual module name
	"log"
	"net/http"
)

const apiBaseURL = "https://groupietrackers.herokuapp.com/api"

// getJSON fetches data from the given URL and unmarshals it into the provided interface.
func getJSON(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error making HTTP request to %s: %v", url, err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("HTTP request to %s failed with status code: %d", url, resp.StatusCode)
		return fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		log.Printf("Error decoding JSON from %s: %v", url, err)
		return err
	}

	return nil
}

// GetArtists fetches the artists data from the API.
func GetArtists(artists *[]models.Artist) error {
	url := apiBaseURL + "/artists"
	return getJSON(url, artists)
}

// GetLocations fetches the locations data from the API.
func GetLocations(locations *models.Locations) error {
	url := apiBaseURL + "/locations"
	return getJSON(url, locations)
}

// GetDates fetches the dates data from the API.
func GetDates(dates *models.Dates) error {
	url := apiBaseURL + "/dates"
	return getJSON(url, dates)
}

// GetRelations fetches the relations data from the API.
func GetRelations(relations *models.Relation) error {
	url := apiBaseURL + "/relation"
	return getJSON(url, relations)
}
