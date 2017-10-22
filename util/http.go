package util

import (
	"fmt"
	"net/http"
)

func Redirect(w http.ResponseWriter, r *http.Request, location string, code int){
	redirect := fmt.Sprintf("%s%s", Fullurl,location)
	http.Redirect(w,r,redirect,code)
}