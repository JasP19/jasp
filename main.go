package main

import (
	"log"
	"os"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// hello is a component that displays a simple "Hello World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type jasp struct {
	app.Compo
}

// The Render method is where the component appearance is defined.
func (j *jasp) Render() app.UI {
	b, err := os.ReadFile("./web/cube.html")
	if err != nil {
		log.Fatal(err)
	}

	return app.Raw(string(b))
}

func main() {
	app.Route("/", &jasp{})
	app.RunWhenOnBrowser()

	err := app.GenerateStaticWebsite(".", &app.Handler{
		Name:        "JasP",
		Description: "A WASM Rubik's cube site",
		Styles: []string{
			"./web/cube.css",
		},
		Scripts: []string{
			"./web/cube.js",
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
