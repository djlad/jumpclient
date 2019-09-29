package jumpserver

import(
	"net/http"
)



func serveMainRoutes(){
	http.HandleFunc("/", mainHandler)
}