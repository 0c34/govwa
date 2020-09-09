package sqli

import(

	"log"
	"strconv"
	"net/http"

	"github.com/julienschmidt/httprouter"

	// /"github.com/govwa/user/session"
	"github.com/govwa/util/middleware"
	"github.com/govwa/util"


)

type SQLI struct{}

func New()SQLI{
	return SQLI{}
}

func (self SQLI)SetRouter(r *httprouter.Router){

	mw := middleware.New()
	r.GET("/sqli1", mw.CapturePanic(mw.AuthCheck(sqli1Handler))) //not use logger due to sqlmap request
	r.GET("/sqli2", mw.CapturePanic(mw.DetectSQLMap(mw.AuthCheck(sqli2Handler))))
}

func sqli1Handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	
	uid := util.GetCookie(r,"Uid")//many developer use this style. set reference key in cookie with no sanitaze

	/*
	this prevent idor injection but not lead to sql injection 	

	s := session.New()
	sid := s.GetSession(r, "id")
	if( sid != uid){
		uid = sid
	} */

	p := NewProfile()

	data := make(map[string]interface{}) //data to send to client

	if(!util.CheckLevel(r)){ //level == low
		err := p.UnsafeQueryGetData(uid)
		if err != nil{
			data["error"] = err.Error()
		}
	}else{
		err := p.SafeQueryGetData(uid)
		if err != nil{
			data["error"] = "No Data Found"
			log.Printf("prepare error : %s", err.Error())
		}
	} 
	data["title"] = "Sql Injection"
	data["uid"] = strconv.Itoa(p.Uid)
	data["name"] = p.Name
	data["city"] = p.City
	data["number"] = p.PhoneNumber

	util.SafeRender(w,r,"template.sqli1",data)

}

func sqli2Handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params){

	uid := r.FormValue("uid")

	p := NewProfile()

	data := make(map[string]interface{}) //data to send to client

	if(!util.CheckLevel(r)){ //level == low
		err := p.UnsafeQueryGetData(uid)
		if err != nil{
			log.Printf("sql error")
		}
	}else{
		err := p.SafeQueryGetData(uid)
		if err != nil{
			data["error"] = "No Data Found"
			log.Printf("prepare error : %s", err.Error())
		}
	}

	data["title"] = "Sql Injection"
	data["name"] = p.Name
	data["city"] = p.City
	data["number"] = p.PhoneNumber
	util.SafeRender(w,r,"template.sqli2",data)

}
