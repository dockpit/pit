package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var bind = flag.String("bind", ":10000", "the address and port on which the token service will bind")
var dbURL = flag.String("db", ":postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full", "database connection string")

func main() {
	flag.Parse()

	log.Printf("Connection to DB '%s'...", *dbURL)
	db, err := sql.Open("postgres", *dbURL)
	if err != nil {
		log.Fatal(err)
	}

	tokens, err := db.Query(`SELECT token FROM tokens`)
	if err != nil {
		log.Fatal(err)
	}

	_ = tokens

	log.Printf("Listening on '%s'...", *bind)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	})

	http.ListenAndServe(*bind, nil)
}
