package main

import (
	"fmt"
	"net/http"
)

import "secureCodingLab/util"
import validation "secureCodingLab/vulnerability/inputvalidation"
import sqli "secureCodingLab/vulnerability/sqli"
import xss "secureCodingLab/vulnerability/xss"

//input validation
func validateHandler(w http.ResponseWriter, r *http.Request) {

	var data = validation.WithNoValidation(r) //default
	if util.CheckLevel(r) {                   //if level == high
		data = validation.WithValidation(r)
	}
	datares := struct {
		Res string
	}{
		Res: data,
	}
	//fmt.Println(checkLevel(r))
	util.SafeRender(w, "validation", datares)
}

//sql injection and escaping
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	data, err := sqli.UnsafeGetData(r) //default function set to unsafe
	if util.CheckLevel(r) { //if level == high
		data, err = sqli.SafeGetData(r)
	}
	if err != nil {
		fmt.Println(err.Error())
	}
	util.RenderAsJson(w, data)
}

//cros site scripting
func getName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	data := xss.New()
	if !util.CheckLevel(r) { //if level == low
		data.UnEscapeString(name)
	}else{
		data.EscapeString(name)
	}
	fmt.Println(data.Name)
	fmt.Fprintf(w, data.Name) //usafe response
}

//index and set cookie
func indexHandler(w http.ResponseWriter, r *http.Request) {
	cookie := util.SetCookieLevel(w, r)
	data := struct {
		Title string
		Level string
	}{
		Title: "Index",
		Level: cookie,
	}
	util.SafeRender(w, "index", data)
}
func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/validate", validateHandler) //input validation
	http.HandleFunc("/getuser", getUserHandler)   //sql injection
	http.HandleFunc("/getinfo", getName)          //cross site scripting

	http.ListenAndServe(":8082", nil)
}
