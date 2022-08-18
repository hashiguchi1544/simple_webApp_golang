package controllers

import (
	"fmt"
	"net/http"
	"text/template"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func StartServer() error {
	files := http.FileServer(http.Dir("app/views"))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/top/", top)

	return http.ListenAndServe(":8080", nil)
}
