package jumpserver

import "fmt"
import "net/http"

func mainHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "serving main")
}

func registerHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "registered handler")
}