package Handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	link "groupie/global"
)

var ApiData link.ApiOfArtist

func PageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorPages(w, 405)
		return
	}
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 || parts[1] != "artists" {
		errorPages(w, 404)
		return
	}

	id := parts[2]
	if id == "" {
		errorPages(w, 404)
		return
	}

	if r.URL.Query().Encode() != "" {
		errorPages(w, 404)
		return
	}

	IdNumber, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || IdNumber <= 0 || IdNumber > 52 {
		errorPages(w, 404)
		return
	}

	// Fetch data from various endpoints
	if err := fetchData(link.Api+"/artists/"+r.PathValue("id"), &ApiData.ArtistData); err != nil {
		errorPages(w, 502)
		return
	}
	if err := fetchData(link.Api+"/locations/"+r.PathValue("id"), &ApiData.Locations); err != nil {
		errorPages(w, 502)
		return
	}
	if err := fetchData(link.Api+"/dates/"+r.PathValue("id"), &ApiData.Dates); err != nil {
		errorPages(w, 502)
		return
	}
	if err := fetchData(link.Api+"/relation/"+r.PathValue("id"), &ApiData.Relation); err != nil {
		errorPages(w, 502)
		return
	}

	// Render template with data
	tmpl, err := template.ParseFiles("templates/result.html")
	if err != nil {
		errorPages(w, 500)
		return
	}
	if err := tmpl.Execute(w, ApiData); err != nil {
		errorPages(w, 500)
		return
	}
}

func fetchData(url string, target interface{}) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return json.NewDecoder(response.Body).Decode(target)
}
