package connection

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func checkResponseCorrectness(resp model.nbplResponse) error {
	if resp.Code != "GBP" || resp.Rates[0].Mid == 0 {
		return errors.New("Error retrieving currency exchange data")
	}
	return nil
}

func getCurrentGBPRate() (*nbplResponse, error) {
	resp, err := http.Get("http://api.nbp.pl/api/exchangerates/rates/a/gbp/lasst/?format=json")
	respCode := resp.StatusCode
	fmt.Println(respCode)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println(respCode)
		fmt.Println(resp.StatusCode)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	jsonBody := &nbplResponse{}
	err = json.Unmarshal(body, jsonBody)

	// fmt.Println(string(body))

	if err != nil {
		return nil, err
	}

	err = checkResponseCorrectness(*jsonBody)
	if err != nil {
		return nil, err
	}

	return jsonBody, err

}
