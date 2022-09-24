package main

import (
	"fmt"
	"net/http"

	"local.packages/rss"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprintf(w, " %s", r.Host)
	} else if r.URL.Path == "/feed" {
		rss.GetRss()
	} else {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":5000", nil)
}
