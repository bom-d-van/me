package main

import (
	"github.com/bom-d-van/me/configs"
	"github.com/codegangsta/martini"
	// "github.com/fvbock/blackfriday"
	"net/http"
)

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})

	m.Get("/thoughts", func() string {
		return "Thoughts!"
	})

	m.Get("/thoughts/:artile_name", func(params martini.Params) string {
		return params["artile_name"]
	})

	println("Serving Me on Port", configs.Port)
	for {
		http.ListenAndServe(configs.Port, m)
	}
}
