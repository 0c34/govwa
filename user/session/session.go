package session

import (
	"log"
	"govwa/util"
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
		HttpOnly: true,
	}

	session.Values["govwa_session"] = true

	err = session.Save(r, w) //safe session and send it to client as cookie

	if err != nil {
		log.Println(err.Error())
	}

	//create new session to store on server side
	if data != nil{
		for key, value := range data {
			session.Values[key] = value
		}
	}

}

func (self *Self) GetSession(r *http.Request, key string) interface{} {
	session, err := store.Get(r, "govwa")

	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return session.Values[key]
}

func (self *Self) DeleteSession(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "govwa")
	if err != nil {
		log.Println(err.Error())
	}

	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	session.Values["govwa_session"] = true
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
