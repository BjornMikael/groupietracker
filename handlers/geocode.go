package handlers

import (
    "encoding/json"
    "groupietracker/utils"
    "net/http"
    "strconv"
)

type Location struct {
	City    string  `json:"city"`
	Country string  `json:"country"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
}

func GeocodeHandler(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("q")

	// Check cache first
	if cached, ok := utils.GetGeocodeCache(location); ok {
		json.NewEncoder(w).Encode(cached)
		return
	}

	// Call Nominatim API
	resp, err := http.Get("https://nominatim.openstreetmap.org/search?format=json&q=" + location)
	if err != nil {
		http.Error(w, "Geocoding failed", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var results []struct {
		Lat string `json:"lat"`
		Lon string `json:"lon"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil || len(results) == 0 {
		http.Error(w, "Location not found", http.StatusNotFound)
		return
	}

	// Cache and return result
	lat, _ := strconv.ParseFloat(results[0].Lat, 64)
	lon, _ := strconv.ParseFloat(results[0].Lon, 64)

	cached := utils.GeocodeResult{Lat: lat, Lon: lon}
	utils.SetGeocodeCache(location, cached)

	json.NewEncoder(w).Encode(cached)
}
