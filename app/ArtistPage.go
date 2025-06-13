package app

import (
	"html/template"
	"net/http"
)

// HomeHandler creates the template that is used to display the page and handles GET request
func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("templates/group-template.html"))
	switch r.Method {
	case "GET":
		if id_int < 0 {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			break
		}
		err := template.Execute(w, Api.Artists[id_int])
		if err != nil {
			ErrorHandler(w, r, http.StatusNotFound)
		}
	default:
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
	}
}
