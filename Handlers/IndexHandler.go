package Handlers

import (
	"encoding/json"
	"net/http"
	"text/template"

	link "groupie/global"
)

type errorType struct {
	ErrorCode string
	Message   string
}

// Render error pages based on the HTTP status code
func errorPages(w http.ResponseWriter, code int) {
	t, err := template.ParseFiles("templates/error.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		t.Execute(w, errorType{ErrorCode: "500", Message: "Internal Server Error."})
		return
	}

	var errorData errorType

	switch code {
	case 404:
		w.WriteHeader(http.StatusNotFound)
		errorData = errorType{ErrorCode: "404", Message: "Sorry, the page you are looking for does not exist."}
	case 405:
		w.WriteHeader(http.StatusMethodNotAllowed)
		errorData = errorType{ErrorCode: "405", Message: "Method not allowed."}
	case 502:
		w.WriteHeader(http.StatusBadGateway)
		errorData = errorType{ErrorCode: "502", Message: "Failed to fetch data from API"}
	default:
		w.WriteHeader(http.StatusInternalServerError)
		errorData = errorType{ErrorCode: "500", Message: "Internal Server Error."}
	}

	if err = t.Execute(w, errorData); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		t.Execute(w, errorType{ErrorCode: "500", Message: "Internal Server Error."})
	}
}

// Handle requests to the index page, displaying all artists
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorPages(w, 405)
		return
	}
	if r.URL.Path != "/" {
		errorPages(w, 404)
		return
	}

	var artistsData []link.ArtistData

	// Fetch artist data from the API
	response, err := http.Get(link.Api + "/artists")
	if err != nil {
		errorPages(w, 502)
		return
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&artistsData)
	if err != nil {
		errorPages(w, 500)
		return
	}

	// Parse and execute the template with the artist data
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		errorPages(w, 500)
		return
	}

	if err := tmpl.Execute(w, artistsData); err != nil {
		errorPages(w, 500)
		return
	}
}
