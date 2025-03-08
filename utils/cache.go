package utils

import (
	"sync"
)

var (
	geocodeCache = make(map[string]GeocodeResult) // Use GeocodeResult from geocode.go
	cacheMutex   = &sync.Mutex{}
)

// GetGeocodeCache retrieves a cached geocode result
func GetGeocodeCache(location string) (GeocodeResult, bool) {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	result, exists := geocodeCache[location]
	return result, exists
}

// SetGeocodeCache stores a geocode result in the cache
func SetGeocodeCache(location string, result GeocodeResult) {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	geocodeCache[location] = result
}
