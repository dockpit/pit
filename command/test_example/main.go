package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"labix.org/v2/mgo"
)

func handler(c web.C, w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users" {

		//call dependency
		_, err := http.Get("http://192.168.59.103:4321/customers")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//simulate state access
		_, err = mgo.DialWithTimeout("mongodb://192.168.59.103:31000", time.Millisecond*100)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"id":1}]`+"\n")
	} else if r.URL.Path == "/users/21" {

		_, err := mgo.DialWithTimeout("mongodb://192.168.59.103:31000", time.Millisecond*100)
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
	goji.Get("/*", handler)
	goji.Serve()
}
