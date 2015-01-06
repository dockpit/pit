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

		fmt.Fprintf(w, `[]`+"\n")
	} else if r.URL.Path == "/users/21" {
		fmt.Fprintf(w, `{"id": 21}`+"\n")
	}
}

//a fake server for test_test.go
func main() {
	goji.Get("/*", handler)
	goji.Serve()
}
