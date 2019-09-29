package jumpserver

import(
	"fmt";
	"net/http";
	"log"
)

func Main(){
	fmt.Println("Serve Routes")
	serveMainRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}