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
	respCode := resp.StatusCode
	fmt.Println(respCode)
	defer resp.Body.Close()

	//place to handle error responses 404 etc
	if err != nil {
		fmt.Println(respCode)
		fmt.Println(resp.StatusCode)
		return nil, err
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

func CalculateSentAmount(sentAmount, rate float64) float64 {
	return sentAmount * rate
}

func CalculateReceivedAmount(receivedAmount, rate float64) (float64, error) {
	if rate == 0 {
		return 0, errors.New("Division by 0 error")
	}

	return receivedAmount / rate, nil

}
