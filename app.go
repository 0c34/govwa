package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"govwa/util"
	"govwa/util/middleware"
	"govwa/user"
	"govwa/user/session"

	"govwa/vulnerability/sqli"
)

const (
	banner = `
     ÛÛÛÛÛÛÛÛÛ           ÛÛÛÛÛ   ÛÛÛÛÛ ÛÛÛÛÛ   ÛÛÛ   ÛÛÛÛÛ   ÛÛÛÛÛÛÛÛÛ  
    ÛÛÛ°°°°°ÛÛÛ         °°ÛÛÛ   °°ÛÛÛ °°ÛÛÛ   °ÛÛÛ  °°ÛÛÛ   ÛÛÛ°°°°°ÛÛÛ 
   ÛÛÛ     °°°   ÛÛÛÛÛÛ  °ÛÛÛ    °ÛÛÛ  °ÛÛÛ   °ÛÛÛ   °ÛÛÛ  °ÛÛÛ    °ÛÛÛ 
  °ÛÛÛ          ÛÛÛ°°ÛÛÛ °ÛÛÛ    °ÛÛÛ  °ÛÛÛ   °ÛÛÛ   °ÛÛÛ  °ÛÛÛÛÛÛÛÛÛÛÛ 
  °ÛÛÛ    ÛÛÛÛÛ°ÛÛÛ °ÛÛÛ °°ÛÛÛ   ÛÛÛ   °°ÛÛÛ  ÛÛÛÛÛ  ÛÛÛ   °ÛÛÛ°°°°°ÛÛÛ 
  °°ÛÛÛ  °°ÛÛÛ °ÛÛÛ °ÛÛÛ  °°°ÛÛÛÛÛ°     °°°ÛÛÛÛÛ°ÛÛÛÛÛ°    °ÛÛÛ    °ÛÛÛ 
   °°ÛÛÛÛÛÛÛÛÛ °°ÛÛÛÛÛÛ     °°ÛÛÛ         °°ÛÛÛ °°ÛÛÛ      ÛÛÛÛÛ   ÛÛÛÛÛ
     °°°°°°°°°   °°°°°°       °°°           °°°   °°°      °°°°°   °°°°° `
)

//index and set cookie

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	cookie := util.SetCookieLevel(w, r)
	data := make(map[string]interface{})
	
	s := session.New()
	uname := s.GetSession(r, "uname")
	id := s.GetSession(r, "id")
	
	data["level"] = cookie
	data["title"] = "Index"
	data["weburl"] = util.Fullurl
	data["uname"] = uname
	data["uid"] = id

	fmt.Println(r.FormValue("govwa_session"))
	util.SafeRender(w,"template.index", data)
}

func main() {
	fmt.Println(banner)
	mw := middleware.New()
	router := httprouter.New()
	userObj := user.New()
	sqlI := sqli.New()

	router.ServeFiles("/public/*filepath", http.Dir("public/"))
	router.GET("/", mw.LoggingMiddleware(mw.AuthCheck(indexHandler)))
	router.GET("/index", mw.LoggingMiddleware(mw.AuthCheck(indexHandler)))

	userObj.SetRouter(router)
	sqlI.SetRouter(router)
	
	s := http.Server{
		Addr : ":8082",
		Handler : router,
	}

	fmt.Printf("Server running at port %s\n", s.Addr)
	s.ListenAndServe()

}
