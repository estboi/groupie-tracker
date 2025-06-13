package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"group/app"
)

func main() {
	// Sets the max time limit to download data from API
	app.Client = &http.Client{Timeout: 10 * time.Second}

	//Downloads new data from API at program start and every 10 minutes
	go func() {
		for {
			app.MakeApi()
			time.Sleep(10 * time.Minute)
		}
	}()
	// Adds files from /static folder to handler so that they can be used
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("assets"))))
	http.Handle("/icons/", http.StripPrefix("/icons", http.FileServer(http.Dir("icons"))))
	http.Handle("/functions/", http.StripPrefix("/functions", http.FileServer(http.Dir("functions"))))
	http.HandleFunc("/", app.HomeHandler)
	http.HandleFunc("/artist", app.ArtistHandler)
	http.HandleFunc("/filter", app.FilterHandler)
	// http.HandleFunc("/location", app.LocationHandler)
	fmt.Printf("Starting server at port 8080\n\n")
	fmt.Printf("http://localhost:8080/\n")

	// Creates the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("%v - internal server error", http.StatusInternalServerError)
	}
}
