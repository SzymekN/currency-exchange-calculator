package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SzymekN/currency-exchange-calculator/pkg/webgui"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func main() {

	app.Route("/", &webgui.CalculatorForm{})

	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Author:      "Szymon Nowak",
		Name:        "Currency Exchange",
		Title:       "Currency Exchange",
		Description: "Currency exchange calculator",
		Icon: app.Icon{
			Default: "/web/icon.png",
		},
		Styles: []string{
			`/web/style.css`,
		},
	})

	fmt.Println("Server listening on localhost:8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
