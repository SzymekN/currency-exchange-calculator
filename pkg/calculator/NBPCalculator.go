package calculator

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/SzymekN/currency-exchange-calculator/model"
)

const gbpDefaultURL = "http://api.nbp.pl/api/exchangerates/rates/a/gbp/last/?format=json"

type HttpGetter interface {
	Get(url string) (resp *http.Response, err error)
}

type DefaultHttpGetter struct {
}

func (d DefaultHttpGetter) Get(url string) (resp *http.Response, err error) {
	return http.Get(url)
}

func checkResponseCorrectness(wantedCurrencyCode string, resp model.NBPResponse) error {

	if resp.Rates == nil {
		return errors.New("No values received")
	}

	if resp.Rates[0].Mid == 0 {
		return errors.New("Invalid exchange value received")
	}

	if resp.Code != wantedCurrencyCode {
		return errors.New("Invalid currency received")
	}

	return nil
}

func makeApiRequest(h HttpGetter, url string) ([]byte, error) {
	resp, err := h.Get(url)
	fmt.Println(resp)
	// fmt.Println(err)

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	// to trigger enter invalid currency code
	if resp.StatusCode == 404 {
		return nil, errors.New("Data not found")
	}

	// to trigger - enter wrong date in url
	if resp.StatusCode == 400 {
		return nil, errors.New("Bad request made")
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil

}

func GetCurrentGBPRate() (float64, error) {
	d := DefaultHttpGetter{}
	body, err := makeApiRequest(d, gbpDefaultURL)

	if err != nil {
		return 0, err
	}

	jsonBody := &model.NBPResponse{}
	err = json.Unmarshal(body, jsonBody)

	// fmt.Println(string(body))

	if err != nil {
		return 0, err
	}

	err = checkResponseCorrectness("GBP", *jsonBody)
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
		return 0, errors.New("Division by 0 error")
	}

	return sentAmount / rate, nil

}
