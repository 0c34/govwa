package middleware

import(
	"log"
	"time"
	"net/http"

	"govwa/user/session"
	"govwa/util"
	"github.com/julienschmidt/httprouter"
)


type Class struct{}

func New()*Class{
	return &Class{}
}

func(self *Class) LoggingMiddleware(h httprouter.Handle) httprouter.Handle{
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
		start := time.Now()
		log.Printf("Request From %s", r.Header.Get("User-Agent"))
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		h(w, r, ps)
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	}
}

func (this *Class) AuthCheck(h httprouter.Handle) httprouter.Handle {
	var sess = session.New()
 	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if !sess.IsLoggedIn(r) {
			redirect := util.Fullurl + "login"
			http.Redirect(w, r, redirect, http.StatusSeeOther)
			return
		}

		h(w, r, ps)
	}
}