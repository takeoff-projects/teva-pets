package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"os"
)


func indexHandler(w http.ResponseWriter, r *http.Request) {

	var tpl = template.Must(template.ParseFiles("templates/index.html", "templates/layout.html"))

	buf := &bytes.Buffer{}
	err := tpl.Execute(buf, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	buf.WriteTo(w)
	log.Println("Home Page Served")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {

	var tpl = template.Must(template.ParseFiles("templates/about.html", "templates/layout.html"))

	buf := &bytes.Buffer{}
	err := tpl.Execute(buf, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	buf.WriteTo(w)
	log.Println("About Page Served")
}
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		
	}
	log.Printf("Port set to: %s", port)

	fs := http.FileServer(http.Dir("assets"))
	mux := http.NewServeMux()

	// This serves the static files in the assets folder 
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// The rest of the routes
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/", indexHandler)

	log.Printf("Webserver listening on Port: %s", port)
	http.ListenAndServe(":"+port, mux)
	
}