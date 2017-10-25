package session

import (

	"govwa/util"
	"log"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

type Self struct{}

func New() *Self {
	return &Self{}
}

var store = sessions.NewCookieStore([]byte(util.Cfg.Sessionkey))

func (self *Self) SetSession(w http.ResponseWriter, r *http.Request, data map[string]string) {
	session, err := store.Get(r, "govwa")

	if err != nil {
		log.Println(err.Error())
	}

	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400,
		HttpOnly: false, //set to false for xss :) 
	}

	session.Values["govwa_session"] = true
	session.Values["asu"] = "coba"
	//create new session to store on server side
	if data != nil {
		for key, value := range data {
			session.Values[key] = value
		}
	}
	err = session.Save(r, w) //safe session and send it to client as cookie
	
		if err != nil {
			log.Println(err.Error())
		}
}

func (self *Self) GetSession(r *http.Request, key string) string  {
	session, err := store.Get(r, "govwa")

	if err != nil {
		log.Println(err.Error())
		return ""
	}
	data := session.Values[key]
	sv := fmt.Sprintf("%v", data)
	return sv
}

func (self *Self) DeleteSession(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "govwa")
	if err != nil {
		log.Println(err.Error())
	}

	session.Values["govwa_session"] = false
	err = session.Save(r, w) //safe session and send it to client as cookie

	if err != nil {
		log.Println(err.Error())
	}

	return
}

func (self *Self) IsLoggedIn(r *http.Request) bool {
	s, err := store.Get(r, "govwa")
	if err != nil {
		log.Println(err.Error())
	}
	if auth, ok := s.Values["govwa_session"].(bool); !ok || !auth {
		return false
	}
	return true
}


/* func init(){
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}

} */