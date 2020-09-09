package xss

import (
	"fmt"
	"log"
	"html"
	"regexp"
	"net/http"
	"html/template"

	"github.com/julienschmidt/httprouter"

	"github.com/govwa/util"
	"github.com/govwa/util/middleware"
	"github.com/govwa/vulnerability/sqli"
)

type XSS struct{
	Name string
}
func New()XSS{
	return XSS{}
}
func (self XSS)SetRouter(r *httprouter.Router){
	mw := middleware.New()
	r.GET("/xss1", mw.LoggingMiddleware(mw.CapturePanic(mw.AuthCheck(xss1Handler))))
	r.POST("/xss1", mw.LoggingMiddleware(mw.CapturePanic(mw.AuthCheck(xss1Handler))))
	r.GET("/xss2", mw.LoggingMiddleware(mw.CapturePanic(mw.AuthCheck(xss2Handler))))
}

func xss1Handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	
	/* template.HTML is a vulnerable function */
	
	data := make(map[string]interface{})

	if r.Method == "GET"{
		
		term := r.FormValue("term")

		if(util.CheckLevel(r)){ //level = high
			term = HTMLEscapeString(term)
		}

		if term == "sql injection"{
			term = "sqli"
		}
		
		term = removeScriptTag(term)
		vulnDetails := GetExp(term)

		notFound := fmt.Sprintf("<b><i>%s</i></b> not found",term)
		value := fmt.Sprintf("%s", term)

		if term == ""{
			data["term"] = ""
		}else if vulnDetails == ""{
			data["value"] = template.HTML(value)
			data["term"] = template.HTML(notFound) //vulnerable function
		}else{
			vuln := fmt.Sprintf("<b>%s</b>",term)
			data["value"] = template.HTML(value)
			data["term"] = template.HTML(vuln)
			data["details"] = vulnDetails
		}

	}
	data["title"] = "Cross Site Scripting"
	util.SafeRender(w,r, "template.xss1", data)
}

func xss2Handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params){

	uid := r.FormValue("uid")

	if(util.CheckLevel(r)){ //level = high
		uid = HTMLEscapeString(uid)
	}

	p := sqli.NewProfile() //using sqli get profile module instead of create new function
	err := p.SafeQueryGetData(uid)

	if err != nil{
		log.Println(err.Error())
	}

	data := make(map[string]interface{})

	js := ` <script>
			var id = %s
			var name = "%s"
			var city = "%s"
			var number = "%s"
			</script>` //here is the mistake, render value to a javascript that came from client request

	inlineJS := fmt.Sprintf(js,uid, p.Name, p.City, p.PhoneNumber)

	data["title"] = "Cross Site Scripting"

	data["inlineJS"] = template.HTML(inlineJS) //this will render the javascript on client browser

	util.SafeRender(w, r, "template.xss2", data)

}

func HTMLEscapeString(text string)string{
	filter := regexp.MustCompile("<[^>]*>")
	output := filter.ReplaceAllString(text,"")
	return html.EscapeString(output)
}

func removeScriptTag(text string)string{
	filter := regexp.MustCompile("<script*>.*</script>")
	output := filter.ReplaceAllString(text,"")
	return output
}