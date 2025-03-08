package handlers

import (
	"groupietracker/models" // Replace "yourmodule" with your actual module name
	"html/template"
	"log"
	"net/http"
)

// Define a struct to hold the data that will be passed to the template
type HomePageData struct {
	Artists []models.Artist
	Title   string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

// HomeHandler displays the list of artists using a template
func HomeHandler(w http.ResponseWriter, r *http.Request, artists []models.Artist) {
	// Create an instance of HomePageData and populate it with the artists data
	data := HomePageData{
		Artists: artists,
		Title:   "Home", // Set the title for the home page
	}

	// Execute the template, passing in the data
	err := tpl.ExecuteTemplate(w, "base.html", data) // Use base.html as the entry point
	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
