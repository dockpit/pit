package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"labix.org/v2/mgo"
)

var mongo = flag.String("mongo", "mongodb://localhost:31000", "test should provide state ip")
var token = flag.String("token", "http://localhost:4321", "test should provide dep ip")

func handler(c web.C, w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users" {

		//call dependency
		_, err := http.Get(fmt.Sprintf("%s/health", *token))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//simulate state access
		_, err = mgo.DialWithTimeout(*mongo, time.Millisecond*100)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"id":1}]`+"\n")
	} else if r.URL.Path == "/users/21" {

		_, err := mgo.DialWithTimeout(*mongo, time.Millisecond*100)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.NotFound(w, r)
		return
	}
}

//a fake server for test_test.go
func main() {
	flag.Parse()

	goji.Get("/*", handler)
	goji.Serve()
}
