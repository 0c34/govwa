package main

import (
	"fmt"
	"net/http"
)
import "secureCodingLab/util"
import validation "secureCodingLab/vulnerability/inputvalidation"
import sqli "secureCodingLab/vulnerability/sqli"

func indexHandler(w http.ResponseWriter, r *http.Request) {

}

//input validation
func validateHandler(w http.ResponseWriter, r *http.Request) {
	data := validation.WithValidation(r)
	datares := struct {
		Res string
	}{
		Res: data,
	}
	fmt.Println(datares.Res)
	util.SafeRender(w, "validation", datares)
}

//sql injection and escaping
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	data, err := sqli.UnsafeGetData(r)
	if err != nil {
		fmt.Println(err.Error())
	}
	util.RenderAsJson(w, data)
}
func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/validate", validateHandler)
	http.HandleFunc("/getuser", getUserHandler)

	http.ListenAndServe(":8082", nil)
}
