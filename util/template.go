package util

import(
	"html/template"
	"encoding/json"
	"net/http"
	"fmt"
	"os"
)

func SafeRender(w http.ResponseWriter, tmpl string, p interface{}) {
	t, err := template.ParseFiles("templates/"+tmpl + ".html")
	if err != nil {
		fmt.Printf(err.Error())
	}
	t.Execute(w, p)
}
func RenderAsJs(w http.ResponseWriter, tmpl string, jsscript string){
	t, err := template.ParseFiles("templates/" + tmpl + ".html")
	if err != nil{
		fmt.Printf(err.Error())
	}
	t.ExecuteTemplate(os.Stdout,"T", jsscript)
}

func RenderAsJson(w http.ResponseWriter, data []interface{})[]byte{
	b, err := json.Marshal(data)
	if err != nil{
		fmt.Println(err.Error())
	}
	return b
}