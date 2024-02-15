package user

import (

	"log"
	"net/http"
	"strconv"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"html/template"

	"github.com/govwa/util"
	"github.com/govwa/util/config"
	"github.com/govwa/user/session"
	"github.com/govwa/util/database"
	"github.com/govwa/util/middleware"

	"github.com/julienschmidt/httprouter"
)

/*
uname : admin
pass : govwaadmin

uname : user1
pass : govwauser1

*/

type Self struct{} //oop like syntax

func New() *Self {
	return &Self{}
}
func (self *Self) SetRouter(r *httprouter.Router) {
	/* register all router */

	mw := middleware.New() //implement middleware

	r.GET("/login", mw.LoggingMiddleware(mw.CapturePanic(LoginViewHandler)))
	r.POST("/login", mw.LoggingMiddleware(mw.CapturePanic(LoginViewHandler)))
	r.GET("/logout", mw.LoggingMiddleware(Logout))
}

func LoginViewHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	/* handler for login view */

	/* check database for setup */
	ok, err := database.CheckDatabase()
	if !ok || err != nil{
		util.Redirect(w, r, "setup", 302) //if no database will redirect to setup page
	}
	
	/* value of data will send to client over template */
	data := make(map[string]interface{})
	data["Title"] = "Login"
	data["govwahost"] = config.Fullurl

	s := session.New()

	if s.IsLoggedIn(r) { //if user session isset wkwk redirect to index page
		util.Redirect(w, r, "index", 302)
	}

	if r.Method == "POST" {

		if !validateForm(w,r,ps) {
			data["message"] = template.HTML("<div id=\"message\" class=\"alert alert-danger\"><p>Empty Username or Password</p></div>")
		}else{
			if loginAction(w, r, ps) {
				util.Redirect(w, r, "index", 302)
			} else {
				//the best solution instead of using ajax request
				data["message"] = template.HTML("<div id=\"message\" class=\"alert alert-danger\"><p>Incorrect Username or Password</p></div>")
				log.Println("Login Failed")
			}
		}
	}
	util.SafeRender(w,r, "template.login", data)
}

func loginAction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) bool {

	/* handler for login action */
	uname := r.FormValue("username")
	pass := Md5Sum(r.FormValue("password"))

	uData := checkUserQuery(uname, pass) //handle user data from db
	if uData.cnt == 1 {
		s := session.New()

		/* save user data to session */
		sessionData := make(map[string]string)
		sessionData["uname"] = uData.uname
		sessionData["id"] = strconv.Itoa(uData.id)

		s.SetSession(w, r, sessionData)
		util.SetCookie(w, "Uid", strconv.Itoa(uData.id)) //save user_id to cookie
		log.Println("Login Success")
		return true
	} else {
		return false
	}
}

func Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	s := session.New()
	s.DeleteSession(w, r)
	cookies := []string{"Level", "Uid"}
	util.DeleteCookie(w,cookies)
	util.Redirect(w, r, "login", 302)
}

func validateForm(w http.ResponseWriter, r *http.Request, _ httprouter.Params)bool{
	uname := r.FormValue("username")
	pass := r.FormValue("password")
	if uname == "" || pass == ""{
		return false
	}
	return true
}

/* type to handle user data that return form query */
type UserData struct {
	id    int
	uname string
	cnt int
}

var db *sql.DB

func checkUserQuery(username, pass string) *UserData {
	/* this function will check rows num which return from query */
	db, err := database.Connect()
	if err != nil {
		log.Println(err.Error())
	}

	var uData = UserData{} //inisialize empty userdata

	const (
		sql = `SELECT id, uname, COUNT(*) as cnt
						FROM Users 
						WHERE uname=? 
						AND pass=?`)

	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Println(err.Error())
	}
	defer stmt.Close()
	err = stmt.QueryRow(username, pass).Scan(&uData.id, &uData.uname, &uData.cnt)
	return &uData

}

func Md5Sum(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
