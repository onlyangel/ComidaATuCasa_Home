package deliveringfood

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var sessionsStore = sessions.NewCookieStore([]byte("L|=ps&1]+-Cox=Q-K53F3'E7GF3\"yQ33P6@%}gf=b75V7\"y6#z120L^!8r%%D8e$[6'X]@A55st;2yE{X!7b6v6~OG/@j1\"4?.Ms"))



func init() {
    http.HandleFunc("/", rootWS)
    //http.HandleFunc("/dashboard", dashboard)

	//register_user_handlers()
}

func rootWS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w,r,"/public/index.html",http.StatusFound)
}

func dashboard(w http.ResponseWriter, r *http.Request) {
	//secureWebAccess(w,r)
	//http.Redirect(w, r, "/dashboard", http.StatusMovedPermanently)
	//fmt.Fprintf(w,"dashboard", http.)
}

