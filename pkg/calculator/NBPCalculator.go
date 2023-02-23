package calculator

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/SzymekN/currency-exchange-calculator/model"
)

const GBPDefaultURL = "http://api.nbp.pl/api/exchangerates/rates/a/gbp/last/?format=json"

var NoValuesReceivedErr = errors.New("No values received")
var InvalidExchangeValueErr = errors.New("Invalid exchange value received")
var InvalidCurrencyReceivedErr = errors.New("Invalid currency received")
var NotFoundErr = errors.New("Not found")
var BadRequestErr = errors.New("Bad request made")
var DivisionBy0Err = errors.New("Division by 0 error")

func checkResponseCorrectness(wantedCurrencyCode string, resp model.NBPResponse) error {

	if resp.Rates == nil {
		return NoValuesReceivedErr
	}

	if resp.Rates[0].Mid == 0 {
		return InvalidExchangeValueErr
	}

	if resp.Code != wantedCurrencyCode {
		return InvalidCurrencyReceivedErr
	}

	return nil
}

func makeApiRequest(h HttpGetter, url string) ([]byte, error) {
	resp, err := h.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// to trigger enter invalid currency code
	if resp.StatusCode == 404 {
		return nil, NotFoundErr
	}

	// to trigger - enter wrong date in url
	if resp.StatusCode == 400 {
		return nil, BadRequestErr
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil

}

func GetCurrentRate(d HttpGetter, currency, url string) (float64, error) {
	body, err := makeApiRequest(d, url)

	if err != nil {
		return 0, err
	}

	jsonBody := &model.NBPResponse{}
	err = json.Unmarshal(body, jsonBody)

	if err != nil {
		return 0, err
	}

	err = checkResponseCorrectness(currency, *jsonBody)
	if err != nil {
		return 0, err
	}

	return jsonBody.Rates[0].Mid, err

}

func CalculateSentAmount(receivedAmount, rate float64) float64 {
	return receivedAmount * rate
}

func CalculateReceivedAmount(sentAmount, rate float64) (float64, error) {
	if rate == 0 {
		return 0, DivisionBy0Err
	}

	return sentAmount / rate, nil

}
