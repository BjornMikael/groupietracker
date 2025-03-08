package utils

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type GeocodeResult struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// GeocodeLocation fetches latitude and longitude for a location using Nominatim API
func GeocodeLocation(location string) (lat, lon float64) {
	// Call Nominatim API
	resp, err := http.Get("https://nominatim.openstreetmap.org/search?format=json&q=" + location)
	if err != nil {
		return 0, 0
	}
	defer resp.Body.Close()

	var results []struct {
		Lat string `json:"lat"`
		Lon string `json:"lon"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil || len(results) == 0 {
		return 0, 0
	}

	// Parse latitude and longitude
	lat, _ = strconv.ParseFloat(results[0].Lat, 64)
	lon, _ = strconv.ParseFloat(results[0].Lon, 64)

	return lat, lon
}
