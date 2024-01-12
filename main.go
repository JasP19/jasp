package main

import (
	"log"
	"os"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type jasp struct {
	app.Compo
}

// The Render method is where the component appearance is defined.
func (j *jasp) Render() app.UI {
	b, err := os.ReadFile("web/cube.html")
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
		ShortName:   "JasP",
		Title:       "JasP",
		Description: "A WASM Rubik's cube site",
		Image:       "./web/logo.png",
		Icon: app.Icon{
			Default: "./web/logo.png", // Specify default favicon.
			Large:   "./web/logo.png", // Specify icon on IOS devices.
			SVG:     "./web/logo.png", // Specify SVG version of favicon.
		},
		Styles: []string{
			"./web/cube.css",
		},
		// manually updated index.html to defer the script
		RawHeaders: []string{
			`<script defer src="./web/cube.js"></script>`,
		},
		// Scripts: []string{
		// 	"./web/cube.js",
		// },
	})

	if err != nil {
		log.Fatal(err)
	}
}
