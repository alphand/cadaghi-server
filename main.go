package main

import (
	"io"
	"log"
	"net/http"
)

func hello(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "Hello World! path: "+req.URL.Path)
}

func main() {
	log.Print("Server is ready at :9090")

	http.HandleFunc("/", hello)
	http.ListenAndServe(":9090", nil)
}
