package calculator

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/SzymekN/currency-exchange-calculator/model"
)

const gbpInvalidCurrencyURL = "http://api.nbp.pl/api/exchangerates/rates/a/gabp/last/?format=json"
const gbpInvalidDateURL = "http://api.nbp.pl/api/exchangerates/rates/a/gbp/3023-01-01/?format=json"

func TestDefaultHttpGetter_Get(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name     string
		d        DefaultHttpGetter
		args     args
		wantResp *http.Response
		wantErr  bool
	}{
		{
			name:     "Positive case",
			d:        DefaultHttpGetter{},
			args:     args{url: gbpDefaultURL},
			wantResp: &http.Response{StatusCode: 200},
			wantErr:  false,
		},
		{
			name:     "404 error test case",
			d:        DefaultHttpGetter{},
			args:     args{url: gbpInvalidCurrencyURL},
			wantResp: &http.Response{StatusCode: 404},
			wantErr:  false,
		},
		{
			name:     "400 error test case",
			d:        DefaultHttpGetter{},
			args:     args{url: gbpInvalidDateURL},
			wantResp: &http.Response{StatusCode: 400},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := tt.d.Get(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("DefaultHttpGetter.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResp.StatusCode != tt.wantResp.StatusCode {
				t.Errorf("DefaultHttpGetter.Get() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func Test_checkResponseCorrectness(t *testing.T) {
	type args struct {
		wantedCurrencyCode string
		resp               model.NBPResponse
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Positive test case",
			args: args{
				wantedCurrencyCode: "GBP",
				resp: model.NBPResponse{
					Code: "GBP",
					Rates: []model.NBPRates{
						{Mid: 5.0},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Wrong currency",
			args: args{
				wantedCurrencyCode: "GBP",
				resp: model.NBPResponse{
					Code: "USD",
					Rates: []model.NBPRates{
						{Mid: 5.0},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "No value in response",
			args: args{
				wantedCurrencyCode: "GBP",
				resp: model.NBPResponse{
					Code: "GBP",
				},
			},
			wantErr: true,
		},
		{
			name: "Exchange rate empty",
			args: args{
				wantedCurrencyCode: "GBP",
				resp: model.NBPResponse{
					Code: "GBP",
					Rates: []model.NBPRates{
						{},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkResponseCorrectness(tt.args.wantedCurrencyCode, tt.args.resp); (err != nil) != tt.wantErr {
				t.Errorf("checkResponseCorrectness() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_makeApiRequest(t *testing.T) {
	type args struct {
		h   HttpGetter
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := makeApiRequest(tt.args.h, tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("makeApiRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeApiRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCurrentGBPRate(t *testing.T) {
	tests := []struct {
		name    string
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCurrentGBPRate()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCurrentGBPRate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetCurrentGBPRate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateSentAmount(t *testing.T) {
	type args struct {
		receivedAmount float64
		rate           float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Positive 25 * 5",
			args: args{receivedAmount: 25, rate: 5},
			want: 125,
		},
		{
			name: "Positive 5 * 1",
			args: args{receivedAmount: 5, rate: 1},
			want: 5,
		},
		{
			name: "Positive 100 * 0.1",
			args: args{receivedAmount: 100, rate: 0.1},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateSentAmount(tt.args.receivedAmount, tt.args.rate); got != tt.want {
				t.Errorf("CalculateSentAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateReceivedAmount(t *testing.T) {
	type args struct {
		sentAmount float64
		rate       float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name:    "Positive 25 / 5",
			args:    args{sentAmount: 25, rate: 5},
			want:    5,
			wantErr: false,
		},
		{
			name:    "Positive 10 / 0.1",
			args:    args{sentAmount: 10, rate: 0.1},
			want:    100,
			wantErr: false,
		},
		{
			name:    "Negative",
			args:    args{sentAmount: 25, rate: 0},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateReceivedAmount(tt.args.sentAmount, tt.args.rate)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateReceivedAmount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CalculateReceivedAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}
