package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/bundgaard/go-shorturl/internal/shorten"
)

var (
	port = flag.String("port", "8080", "Port for HTTP listener, default 8080")
)

func main() {
	flag.Parse()
	root := http.NewServeMux()
	api, err := shorten.NewShorten()
	if err != nil {
		log.Fatal(err)
	}
	root.Handle("/", api)
	log.Printf("listening on port %s", *port)
	if err := http.ListenAndServe(":"+*port, root); err != nil {
		log.Fatal(err)
	}
}
