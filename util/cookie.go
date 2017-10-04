
package util

import "net/http"

func SetCookieLevel(w http.ResponseWriter, r *http.Request)string{
	ck := r.FormValue("level")
	level := ck
	if level == ""{
		level = "low"
	}
	cookie := http.Cookie{Name:"level", Value:level}
	http.SetCookie(w,&cookie)
	return level
}

func CheckLevel(r *http.Request) bool{
	level, _:= r.Cookie("level")
	if level.Value == "" || level.Value == "low"{
		return false //set default level to low
	}else if level.Value == "high"{
		return true //level == high
	}else{
		return false // level == low
	}
}