package model

type NbplRates struct {
	No            string
	EffectiveDate string
	Mid           float64
}

type NbplResponse struct {
	Table    string
	Currency string
	Code     string
	Rates    []NbplRates
}
