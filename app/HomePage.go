package app

import (
	"html/template"
	"net/http"
	"strconv"
)

var id_int int
var sorted = true

type errdata struct {
	Message string
}

// HomeHandler creates the template that is used to display the page and handles GET request
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("templates/template.html"))
	if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	Api2.Artists = nil
	Api2.NoMatches = false
	switch r.Method {
	case "GET":
		err := template.Execute(w, Api)
		if err != nil {
			ErrorHandler(w, r, http.StatusNotFound)
		}
	case "POST":
		Filter(w, r)
		SearchButton = r.FormValue("SearchButton")
		if PressedReverse(w, r) {
			Sort(w, r)
			template.Execute(w, Api)
		} else if PressedReverse(w, r) && !PressedNormal(w, r) {
			template.Execute(w, Api)
		} else if PressedNormal(w, r) {
			Sort(w, r)
			template.Execute(w, Api)
		} else if FilteredButton != "" {
			http.Redirect(w, r, "/filter", http.StatusSeeOther)
		} else if SearchButton != "" {
			SearchReciever(w, r)
			http.Redirect(w, r, "/filter", http.StatusSeeOther)
		} else {
			Id_str := r.FormValue("MoreButton")
			id_int, _ = strconv.Atoi(Id_str)
			http.Redirect(w, r, "/artist", http.StatusSeeOther)
		}

	default:
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
	}
}

// ErrorHandler
func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	temp := template.Must(template.ParseFiles("templates/error.html"))
	var er errdata
	switch status {
	case http.StatusNotFound:
		er.Message = "404 Page not found"
	case http.StatusMethodNotAllowed:
		er.Message = "405 Method not allowed"
	case http.StatusBadRequest:
		er.Message = "400 Bad request"
	default:
		er.Message = "Something went wrong"
	}
	temp.Execute(w, er)
}
