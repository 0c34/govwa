package util

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func SafeRender(w http.ResponseWriter, tmpl string, p interface{}) {
	t, err := template.ParseFiles("templates/" + tmpl + ".html")
	if err != nil {
		fmt.Printf(err.Error())
	}
	t.Execute(w, p)
}

func RenderAsJson(w http.ResponseWriter, data ...interface{}) {
	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(b)
	return
}
