package model

type NBPResponse struct {
	Table    string
	Currency string
	Code     string
	Rates    []NBPRates
}
