package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/SzymekN/currency-exchange-calculator/pkg/calculator"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type calculatorForm struct {
	app.Compo

	PLN     string
	GBP     string
	Rate    float64
	RateStr string
}

func (cf *calculatorForm) Render() app.UI {
	cf.updateRate()
	return app.Div().ID("wrapper").Body(
		app.Div().ID("calc_form").Body(
			app.Form().Body(
				app.Label().
					Text("You send"),
				app.Input().
					Class("amountInput").
					ID("receivedInput").
					Type("text").
					Value(cf.GBP).
					Placeholder("0.00").
					// OnFocus(cf.updateRate).
					OnInput(func(ctx app.Context, e app.Event) {
						if v := parseInput(cf, ctx); v != 0 {
							cf.PLN = formatString(v * 2)
						}
					}).
					OnChange(cf.ValueTo(&cf.GBP)).
					On("focusout", func(ctx app.Context, e app.Event) {
						if v := parseInput(cf, ctx); v != 0 {
							cf.GBP = formatString(v)
						}
					}),
				app.Input().
					Class("disabledInput").
					Disabled(true).
					Value("GBP"),
				app.Img().
					Alt("UK flag").
					Src(`/web/uk_flag.png`),
				app.Label().
					Text("They receive").
					Style("clear", "both"),
				app.Input().
					Class("amountInput").
					ID("sentInput").
					Type("text").
					Value(cf.PLN).
					Placeholder("0.00").
					// OnFocus(cf.updateRate).
					OnInput(func(ctx app.Context, e app.Event) {
						if v := parseInput(cf, ctx); v != 0 {
							cf.GBP = formatString(v / 2)
						}
					}).
					OnChange(cf.ValueTo(&cf.PLN)).
					On("focusout", func(ctx app.Context, e app.Event) {
						if v := parseInput(cf, ctx); v != 0 {
							cf.PLN = formatString(v)
						}
					}),
				app.Input().
					Class("disabledInput").
					Disabled(true).
					Value("PLN"),
				app.Img().
					Alt("PL flag").
					Src(`/web/pol_flag.png`),
				app.Label().
					Text("1PLN = "),
				app.Label().
					Text(cf.RateStr).
					Style("font-weight", "bold"),
			),
		),
	)
}

func (cf *calculatorForm) updateRate() {
	v, err := calculator.GetCurrentRate(calculator.DefaultHttpGetter{}, "GBP", calculator.GBPDefaultURL)

	if err != nil || v == 0 {
		cf.RateStr = "0.00"
		cf.Rate = 0
		return
	}

	cf.Rate = v
	cf.RateStr = formatString(v)

}

func formatString(val float64) string {
	return strconv.FormatFloat(val, 'f', 2, 64)
}

func parseInput(cf *calculatorForm, ctx app.Context) float64 {
	if tmpVal, err := strconv.ParseFloat(ctx.JSSrc().Get("value").String(), 64); err != nil || tmpVal < 0 {
		fmt.Println(err)
		cf.GBP = "0.00"
		cf.PLN = "0.00"
		return 0
	} else {
		return tmpVal
	}

}

func (cf *calculatorForm) countReceived(ctx app.Context, e app.Event) {

	if v := parseInput(cf, ctx); v != 0 {
		cf.PLN = formatString(v * 2)
	}

}

func (cf *calculatorForm) countSent(ctx app.Context, e app.Event) {

	if v := parseInput(cf, ctx); v != 0 {
		cf.GBP = formatString(v / 2)
	}

}

func main() {

	app.Route("/", &calculatorForm{})

	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
		Styles: []string{
			`/web/style.css`,
		},
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
