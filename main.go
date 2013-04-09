package main

import (
	"net/http"
	"net/url"
	"log"
	"flag"
	"fmt"
)

var port int
var root string

func init() {
	flag.IntVar(&port, "p", 8080, "port")
	flag.StringVar(&root, "h", "", "home")
	flag.Parse()

	if len(root) <= 0 {
		log.Fatal("invalid home path")
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path, err := url.QueryUnescape(r.RequestURI)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.ServeFile(w, r, root + path)
	})

	log.Printf("start listen at :%d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",port), nil))
}
