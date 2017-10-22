package user

import (
	"log"
	"fmt"
	"net/http"
	"crypto/md5"
	"encoding/hex"
	"html/template"
	"database/sql"

	"govwa/util"
	"govwa/user/session"
	"govwa/util/middleware"
	"govwa/util/database"

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

	r.GET("/login", mw.LoggingMiddleware(LoginViewHandler))
	r.POST("/login", mw.LoggingMiddleware(LoginViewHandler))
	r.GET("/logout", mw.LoggingMiddleware(Logout))
}

func LoginViewHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	/* handler for login view */

	/* value of data will send to client over template */
	data := make(map[string]interface{})
	data["Title"] = "Login"
	data["govwahost"] = util.Fullurl

	s := session.New()

	if s.IsLoggedIn(r) { //if user session isset wkwk redirect to index page
		util.Redirect(w, r, "index", 302)
	}

	if r.Method == "POST" {
		if loginAction(w, r, ps){
			util.Redirect(w, r, "index", 302)
		}else{
			//the best solution instead of using ajax request
			data["message"] = template.HTML("<div id=\"message\" class=\"alert alert-danger\"><p>Incorrect Username or Password</p></div>")
			log.Println("Login Failed")
		}
	}
	util.SafeRender(w, "template.login", data)
}

func loginAction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) bool{

	/* handler for login action */
	uname := r.FormValue("username")
	pass := Md5Sum(r.FormValue("password"))
	if checkUserQuery(uname, pass) == 1 {
		s := session.New()
		s.SetSession(w, r, nil)
		log.Println("Login Success")
		return true
	} else {
		return false
	}
}

func Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	s := session.New()
	s.DeleteSession(w, r)
	util.Redirect(w, r, "login", 302)
}

var db *sql.DB
func checkUserQuery(username,pass string)int{
	/* this function will check rows num which return from query */
	db, err := database.Connect()
	if err != nil{
		log.Println(err.Error())
	}

	var count int

	sql := fmt.Sprintf(`SELECT COUNT(*) 
						FROM Users 
						WHERE uname=? 
						AND pass=?`)

	stmt,err := db.Prepare(sql)
	if err != nil{
		fmt.Println(err.Error())
	}
	defer stmt.Close()
	err = stmt.QueryRow(username,pass).Scan(&count)
	return count

}

func Md5Sum(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}