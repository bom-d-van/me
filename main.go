package main

import (
	"github.com/codegangsta/martini"
	"net/http"
)

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	for {
		http.ListenAndServe(":80", m)
	}
	// m.Run()
}
