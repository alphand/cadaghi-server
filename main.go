package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	acc "github.com/alphand/skilltree-server/accounts"
)

func helloFunc(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "Hello World! path: "+req.URL.Path)
}

func main() {
	route := mux.NewRouter()
	route.HandleFunc("/", helloFunc)
	route.HandleFunc("/accounts/github", acc.HandleGithubExchange).Methods("POST")

	n := negroni.Classic()
	n.UseHandler(route)
	log.Print("Server is ready at :9090")

	http.ListenAndServe(":9090", n)
}
