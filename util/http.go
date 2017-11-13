package util

import (
	"fmt"
	"net/http"

	"govwa/util/config"
)

func Redirect(w http.ResponseWriter, r *http.Request, location string, code int){
	redirect := fmt.Sprintf("%s%s", config.Fullurl,location)
	http.Redirect(w,r,redirect,code)
}