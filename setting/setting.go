package setting

import(
	
	"fmt"
	"strings"
	"runtime"
	"net/http"

	"github.com/govwa/util"
	"github.com/govwa/user/session"
	"github.com/govwa/util/database"
	"github.com/govwa/util/middleware"

	"github.com/julienschmidt/httprouter"
)

type Setting struct{}

func New() Setting {
	return Setting{}
}

func (self Setting) SetRouter(r *httprouter.Router) {

	mw := middleware.New()

	r.GET("/setting", mw.LoggingMiddleware(mw.CapturePanic(mw.AuthCheck(settingViewHandler))))
	r.POST("/setlevel", mw.LoggingMiddleware(mw.CapturePanic(mw.AuthCheck(setLevelAction))))

}

func settingViewHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	data := make(map[string]interface{})

	data["title"] = "Setting & system info"
	data["dbversion"] = getDBMSVersion()
	data["osversion"] = getOSVersion()
	data["currentuser"] = getCurrentUserName(r)
	data["level"] = getCurrentLevel(r)

	util.SafeRender(w,r,"template.setting", data)
	
}

func setLevelAction(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	if r.Method == "POST" && r.FormValue("level") != ""{
		level := r.FormValue("level")
		util.SetCookieLevel(w, r, level)
		res := struct{
			Res string
		}{
			Res : "Success",
		}
		util.RenderAsJson(w, res)
	}
}

func getDBMSVersion()string{
	const sql=`
			select @@version
	`
	var version string
	db,_ := database.Connect()

	_ = db.QueryRow(sql).Scan(&version)
	dbversion := strings.Split(version, "-")
	return dbversion[0]
}

func getOSVersion()string{
	
	switch os := runtime.GOOS; os{
	case "darwin":
		return "OS X"
	case "linux":
		return "Linux"
	default:
		fmt.Printf("%s.", os)
	}
	return ""
}

func getCurrentUserName(r *http.Request)string{
	s := session.New()
	uname := s.GetSession(r, "uname")
	return uname
}

func getCurrentLevel(r *http.Request)string{
	level := util.GetCookie(r,"Level")
	return level
}