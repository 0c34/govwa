package setup

import(
	"fmt"
	"net/http"

	"govwa/util"
	"govwa/util/database"
	"govwa/util/middleware"

	"github.com/julienschmidt/httprouter"
)

type Setup struct{}

func New() Setup {
	return Setup{}
}

func (self Setup) SetRouter(r *httprouter.Router) {

	mw := middleware.New()

	r.GET("/setup", mw.LoggingMiddleware(mw.CapturePanic(setupViewHandler)))
	r.POST("/setupaction", mw.LoggingMiddleware(mw.CapturePanic(setupActionHandler)))

}

func setupViewHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params){

	var info string
	
	data := make(map[string]interface{})
	ok, err := database.CheckDatabase()
	
	if !ok || err != nil{
		info = fmt.Sprintf(`<div id="info" class="alert alert-danger">%s</div>`,err.Error())
		data["error"] = util.ToHTML(info)
	}

	info = fmt.Sprintf(`<div id="info" class="alert alert-success">Connection Success</div>`)
	data["error"] = util.ToHTML(info)

	data["title"] = "Setup/Reset"

	util.SafeRender(w,r,"template.setup", data)
}

func setupActionHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	
}