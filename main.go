package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/akosourov/forum/api"
	"github.com/akosourov/forum/data"
)

func index(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/content.html",
		"templates/navbar.html",
		"templates/base.html",
	}
	tmpl := template.Must(template.ParseFiles(files...))
	threads := data.ThreadList()
	tmpl.ExecuteTemplate(w, "base", threads)
}

func signup(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/signup.html"))
	switch r.Method {
	case "GET":
		tmpl.Execute(w, nil)
	case "POST":
		email := r.FormValue("email")
		_, err := data.UserByEmail(email)
		if err == nil {
			// success
			// http.Redirect(w, r, "/", http.StatusSeeOther)
			// fmt.Fprint(w, "Ok")
		} else {
			// error msg
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "User already exists")
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func main() {
	fsHdl := http.FileServer(http.Dir("static"))
	mux := http.NewServeMux()

	// web
	mux.Handle("/static/", http.StripPrefix("/static/", fsHdl))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/signup", signup)

	// api
	mux.HandleFunc("/api/v1/threads", api.Threads)

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Error", err)
	}
}
