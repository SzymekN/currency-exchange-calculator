package webgui

import (
	"fmt"
	"strconv"

	"github.com/SzymekN/currency-exchange-calculator/pkg/calculator"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type CalculatorForm struct {
	app.Compo

	PLN     string
	GBP     string
	RateStr string
	Rate    float64
}

func (cf *CalculatorForm) Render() app.UI {
	cf.updateRate()
	return app.Div().ID("wrapper").Body(
		app.Div().ID("calc_form").Body(
			app.Form().Body(
				app.Label().
					Class("inputLabel").
					Text("You send"),
				app.Input().
					Class("amountInput").
					Type("number").
					Value(cf.GBP).
					Placeholder("0,00").
					OnFocus(func(ctx app.Context, e app.Event) {
						cf.updateRate()
					}).
					OnInput(func(ctx app.Context, e app.Event) {
						if v := parseInput(cf, ctx); v != 0 {
							v := calculator.CalculateReceivedAmount(v, cf.Rate)
							cf.PLN = formatInput(v, 2)
						}
					}).
					OnChange(cf.ValueTo(&cf.GBP)).
					On("focusout", func(ctx app.Context, e app.Event) {
						if v := parseInput(cf, ctx); v != 0 {
							cf.GBP = formatInput(v, 2)
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
					Class("inputLabel").
					Text("They receive"),
				app.Input().
					Class("amountInput").
					Type("number").
					Value(cf.PLN).
					Placeholder("0,00").
					OnFocus(func(ctx app.Context, e app.Event) {
						cf.updateRate()
					}).
					OnInput(func(ctx app.Context, e app.Event) {
						if v := parseInput(cf, ctx); v != 0 {
							v, err := calculator.CalculateSentAmount(v, cf.Rate)
							if err != nil {
								cf.GBP = "0,00"
								cf.RateStr = "err"
							} else {
								cf.GBP = formatInput(v, 2)
							}
						}
					}).
					OnChange(cf.ValueTo(&cf.PLN)).
					On("focusout", func(ctx app.Context, e app.Event) {
						if v := parseInput(cf, ctx); v != 0 {
							cf.PLN = formatInput(v, 2)
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
					Text("1GBP = "),
				app.Label().
					ID("rateLabel").
					Text(cf.RateStr),
				app.Label().ID("noFeeLabel").
					Text("No transfer fee"),
			),
		),
	)
}

func (cf *CalculatorForm) updateRate() {
	v, err := calculator.GetCurrentRate(calculator.DefaultHttpGetter{}, "GBP", calculator.GBPDefaultURL)

	if err != nil || v == 0 {
		cf.Rate = 0
		cf.RateStr = "Error while updating exchange rate"
		return
	}

	cf.Rate = v
	cf.RateStr = " " + formatInput(v, 4) + "PLN"

}

func formatInput(val float64, digits int) string {
	return strconv.FormatFloat(val, 'f', digits, 64)
}

func parseInput(cf *CalculatorForm, ctx app.Context) float64 {
	if tmpVal, err := strconv.ParseFloat(ctx.JSSrc().Get("value").String(), 64); err != nil || tmpVal < 0 {
		fmt.Println(err)
		cf.GBP = "0,00"
		cf.PLN = "0,00"
		return 0
	} else {
		return tmpVal
	}

}
