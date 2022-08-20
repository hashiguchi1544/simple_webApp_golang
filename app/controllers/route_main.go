package controllers

import "net/http"

func index(w http.ResponseWriter, r *http.Request) {
	paramsLists := []string{"", "リスト1", "リスト2", "リスト3", "リスト4"}

	generateHTML(w, paramsLists, "layout", "index")
}

func top(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "hello", "layout", "top")
}
