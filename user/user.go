
package user

import(
	"govwa/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Self struct{}

func New()*Self{
	return &Self{}
}
func (self *Self)SetRouter(r *httprouter.Router){
	r.GET("/login", LoginViewHandler)
}

func LoginViewHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	data := make(map[string]interface{})
	data["Title"] = "Login"
	util.SafeRender(w, "template.login", data)
}
