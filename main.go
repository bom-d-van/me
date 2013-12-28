package main

import (
	"github.com/bom-d-van/me/app"
	"github.com/bom-d-van/me/configs"
	"github.com/codegangsta/martini"
	"net/http"
	"os"
)

func main() {
	m := martini.Classic()

	m.Get("/", app.GetThoughts)
	m.Get("/about", app.GetAbout)
	m.Get("/thoughts", app.GetThoughts)
	m.Get("/thoughts/:artile_name", app.GetArticle)
	m.Use(martini.Static(os.Getenv("GOPATH") + "/src/github.com/bom-d-van/me/thoughts"))

	println("Serving Me on Port", configs.Port)
	for {
		http.ListenAndServe(configs.Port, m)
	}
}
