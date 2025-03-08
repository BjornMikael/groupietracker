package utils

import (
	"strconv"
	"strings"
)

// FormatDate formats a date string from "DD-MM-YYYY" to "D.M.YYYY" or "DD.MM.YYYY"
func FormatDate(date string) string {
	parts := strings.Split(date, "-")
	if len(parts) != 3 {
		return date // Return original if invalid format
	}

	// Remove leading zeros
	day, _ := strconv.Atoi(parts[0])
	month, _ := strconv.Atoi(parts[1])
	year := parts[2]

	return strconv.Itoa(day) + "." + strconv.Itoa(month) + "." + year
}

// FormatLocation splits a location string into city and country, and formats them
func FormatLocation(location string) (city, country string) {
	parts := strings.SplitN(location, "-", 2)
	if len(parts) != 2 {
		return location, "" // Fallback if format is invalid
	}

	city = strings.ReplaceAll(parts[0], "_", " ")
	city = strings.Title(city) // Capitalize city name

	country = strings.ReplaceAll(parts[1], "_", " ")
	country = strings.ToUpper(country) // Uppercase country name

	return city, country
}
