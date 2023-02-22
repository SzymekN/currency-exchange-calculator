package model

type nbplRates struct {
	No            string
	EffectiveDate string
	Mid           float64
}

type nbplResponse struct {
	Table    string
	Currency string
	Code     string
	Rates    []nbplRates
}
