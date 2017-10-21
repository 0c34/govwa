package session

import(
	"net/http"
	"fmt"
	"govwa/util"

	"github.com/gorilla/sessions"
)

type Self struct{}

func New()*Self{
	return &Self{}
}

var store = sessions.NewCookieStore([]byte(util.Cfg.Sessionkey))

func (self *Self)SetSession(w http.ResponseWriter, r *http.Request, data map[string]string){
	session, err := store.Get(r, "govwa")
	
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400,
		HttpOnly:true,
	}

	if err != nil{
		fmt.Println(err.Error());
	}

	session.Values["govwa_session"] = true

	err = session.Save(r,w) //safe session and send it to client as cookie

	//create new session to store on server side

	for key,value := range data{
		session.Values[key] = value
	}

	if err != nil{
		fmt.Println(err.Error())
	}
}

func (self *Self) IsLoggedIn(r *http.Request)bool{
	s, err := store.Get(r, "govwa")
	if err != nil{
		fmt.Println(err.Error())
	}
	if auth, ok := s.Values["govwa_session"].(bool); !ok || !auth {
		return false
	}
	return true
}
