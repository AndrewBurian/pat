package main

import (
	"io"
	"log"
	"net/http"

	"github.com/andrewburian/pat"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	if name, err := pat.GetPathParam(req, ":name"); err == nil {
		io.WriteString(w, "context: hello, "+name+"!\n")
	} else {
		io.WriteString(w, "query: hello there")
	}

}

func main() {
	m := pat.New()
	m.UseContext = true // store path params in req.context for efficiency
	m.Get("/hello/:name", http.HandlerFunc(HelloServer))

	// Register this pat with the default serve mux so that other packages
	// may also be exported. (i.e. /debug/pprof/*)
	http.Handle("/", m)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
