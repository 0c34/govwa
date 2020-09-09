package idor

import (

	"log"
	"strconv"
	"net/http"
	"crypto/md5"
	"encoding/hex"

	"github.com/julienschmidt/httprouter"

	"github.com/govwa/user/session"
	"github.com/govwa/util"
	"github.com/govwa/util/middleware"
)

type IDOR struct{}

func New() IDOR {
	return IDOR{}
}

func (self IDOR) SetRouter(r *httprouter.Router) {

	mw := middleware.New()
	r.GET("/idor1", mw.LoggingMiddleware(mw.CapturePanic(mw.AuthCheck(idor1Handler))))
	r.POST("/idor1action", mw.LoggingMiddleware(mw.CapturePanic(mw.AuthCheck(idor1ActionHandler))))
	r.GET("/idor2", mw.LoggingMiddleware(mw.CapturePanic(mw.AuthCheck(idor2Handler))))
	r.POST("/idor2action", mw.LoggingMiddleware(mw.CapturePanic(mw.AuthCheck(idor2ActionHandler))))
}

type DataResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func idor1Handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	session := session.New()
	sid := session.GetSession(r, "id")
	p := NewProfile()
	p.GetData(sid)

	data := make(map[string]interface{})

	data["title"] = "Insecure Direc Object References"
	data["uid"] = strconv.Itoa(p.Uid)
	data["name"] = p.Name
	data["city"] = p.City
	data["number"] = p.PhoneNumber

	util.SafeRender(w, r, "template.idor1", data)

}

func idor2Handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	session := session.New()
	sid := session.GetSession(r, "id")
	p := NewProfile()
	p.GetData(sid)

	data := make(map[string]interface{})
	signature := Md5Sum(sid)

	data["signature"] = signature
	data["title"] = "Insecure Direc Object References"
	data["uid"] = strconv.Itoa(p.Uid)
	data["name"] = p.Name
	data["city"] = p.City
	data["number"] = p.PhoneNumber

	util.SafeRender(w, r, "template.idor2", data)

}

func idor1ActionHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	session := session.New()
	sid := session.GetSession(r, "id")
	p := NewProfile()
	p.GetData(sid)

	/* handle request response with json */
	if r.Method == "POST" {

		cid := util.GetCookie(r, "Uid")
		uid := HTMLEscapeString(r.FormValue("uid"))
		name := HTMLEscapeString(r.FormValue("name"))
		city := HTMLEscapeString(r.FormValue("city"))
		number := HTMLEscapeString(r.FormValue("number"))

		res := &DataResponse{}
		if uid != cid || uid == "" || cid == "" {

			res.Status = "0"
			res.Message = "Missing User Id"
			log.Println("Update Error")

		} else {

			if util.CheckLevel(r) { //level == high
				uid = sid //set uid that fetch from session this use to prevent unauthorize users force update other user profile
			}

			err = p.UpdateProfile(name, city, number, uid)
			if err != nil {
				log.Println(err.Error())
			}
			res.Status = "1"
			res.Message = "Update Success"
			log.Println("Update Success")

		}
		util.RenderAsJson(w, res)
	}
}

func idor2ActionHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	session := session.New()
	sid := session.GetSession(r, "id")
	p := NewProfile()
	p.GetData(sid)

	/* handle request response with json */
	if r.Method == "POST" {
		
		sign := HTMLEscapeString(r.FormValue("signature"))
		uid := HTMLEscapeString(r.FormValue("uid"))
		name := HTMLEscapeString(r.FormValue("name"))
		city := HTMLEscapeString(r.FormValue("city"))
		number := HTMLEscapeString(r.FormValue("number"))

		signature := Md5Sum(uid)

		res := &DataResponse{}
		if sign != signature{

			res.Status = "0"
			res.Message = "Integrity Error"
			log.Println("Update Error")

		} else {

			if util.CheckLevel(r) { //level == high
				uid = sid //set uid that fetch from session this use to prevent unauthorize users force update other user profile
			}

			err = p.UpdateProfile(name, city, number, uid)
			if err != nil {
				log.Println(err.Error())
			}
			res.Status = "1"
			res.Message = "Update Success"
			log.Println("Update Success")

		}
		util.RenderAsJson(w, res)
	}
}

func Md5Sum(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}