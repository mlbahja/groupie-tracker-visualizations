package main

import (
	"fmt"
	"html/template"
	"net/http"

	"groupie/Handlers"
)

type errorType struct {
	ErrorCode string
	Message   string
}

func errorPages(w http.ResponseWriter, code int) {
	t, err := template.ParseFiles("templates/error.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		t.Execute(w, errorType{ErrorCode: "500", Message: "Internal Server Error."})
		return
	} else if code == 404 {
		w.WriteHeader(http.StatusNotFound)
		err = t.Execute(w, errorType{ErrorCode: "404", Message: "Sorry, the page you are looking for does not exist."})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			t.Execute(w, errorType{ErrorCode: "500", Message: "Internal Server Error."})
		}
	} else if code == 405 {
		w.WriteHeader(http.StatusMethodNotAllowed)
		err = t.Execute(w, errorType{ErrorCode: "405", Message: "Method not allowed."})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			t.Execute(w, errorType{ErrorCode: "500", Message: "Internal Server Error."})
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		t.Execute(w, errorType{ErrorCode: "500", Message: "Internal Server Error."})
	}
}

func serveCSS(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/assets/" {
		errorPages(w, 404)
		return
	}
	fs := http.FileServer(http.Dir("./style"))
	http.StripPrefix("/style/", fs).ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/", Handlers.IndexHandler)
	http.HandleFunc("/artists/", Handlers.PageHandler)

	// Serve static files from the "assets" directory
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	fmt.Println("\033[32mServer started at http://127.0.0.1:8080\033[0m")
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
