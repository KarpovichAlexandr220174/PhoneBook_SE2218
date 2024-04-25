package main

import "net/http"

func application(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "main-page.tmpl", nil)
}
