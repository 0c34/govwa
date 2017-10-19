package util

import (
	"encoding/json"
	"html/template"
	"net/http"
)

func SafeRender(w http.ResponseWriter, name string, data map[string]interface{}) {

	template := template.Must(template.ParseGlob("templates/*"))
	template.ExecuteTemplate(w, name, data)
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

func UnSafeRender(w http.ResponseWriter, name string, data ...interface{}) {

	template := template.Must(template.ParseGlob("templates/*"))
	template.ExecuteTemplate(w, name, data)
}
