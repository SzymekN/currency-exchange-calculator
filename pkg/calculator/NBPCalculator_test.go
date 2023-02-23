package calculator

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"

	"github.com/SzymekN/currency-exchange-calculator/model"
	mock_calculator "github.com/SzymekN/currency-exchange-calculator/pkg/calculator/mocks"
	"github.com/SzymekN/currency-exchange-calculator/pkg/fake"
	"github.com/golang/mock/gomock"
)

const gbpInvalidCurrencyURL = "http://api.nbp.pl/api/exchangerates/rates/a/gabp/last/?format=json"
const gbpInvalidDateURL = "http://api.nbp.pl/api/exchangerates/rates/a/gbp/3023-01-01/?format=json"
const InvalidURL = ""

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

func Test_makeApiRequest_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockCalculator := mock_calculator.NewMockHttpGetter(ctrl)
	r := &http.Response{StatusCode: 200, Body: http.NoBody}

	mockCalculator.EXPECT().Get(GBPDefaultURL).Return(r, nil)

	b, err := makeApiRequest(mockCalculator, GBPDefaultURL)

	if b == nil {
		t.Logf("want: []byte, got: nil")
		t.Fail()
	}

	if err != nil {
		t.FailNow()
	}
}
func Test_makeApiRequest_InvalidURL(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockCalculator := mock_calculator.NewMockHttpGetter(ctrl)

	mockCalculator.EXPECT().Get(InvalidURL).Return(nil, errors.New("empty url"))

	b, err := makeApiRequest(mockCalculator, InvalidURL)

	if b != nil {
		t.Logf("want: nil, got: %v\n", b)
		t.Fail()
	}

	if err == nil {
		t.Logf("want: err, got: nil")
		t.FailNow()
	}
}
func Test_makeApiRequest_404Error(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockCalculator := mock_calculator.NewMockHttpGetter(ctrl)
	r := &http.Response{StatusCode: 404, Body: http.NoBody}
	mockCalculator.EXPECT().Get(gbpInvalidCurrencyURL).Return(r, nil)

	b, err := makeApiRequest(mockCalculator, gbpInvalidCurrencyURL)

	if b != nil {
		t.Logf("want: nil, got: %v\n", b)
		t.Fail()
	}

	if !errors.Is(err, NotFoundErr) {
		t.FailNow()
	}
}
func Test_makeApiRequest_400Error(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockCalculator := mock_calculator.NewMockHttpGetter(ctrl)
	r := &http.Response{StatusCode: 400, Body: http.NoBody}
	mockCalculator.EXPECT().Get(gbpInvalidDateURL).Return(r, nil)

	b, err := makeApiRequest(mockCalculator, gbpInvalidDateURL)

	if b != nil {
		t.Logf("want: nil, got: %v\n", b)
		t.Fail()
	}

	if !errors.Is(err, BadRequestErr) {
		t.FailNow()
	}
}
func Test_makeApiRequest_respBodyError(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockCalculator := mock_calculator.NewMockHttpGetter(ctrl)
	r := &http.Response{StatusCode: 200, Body: fake.ReadCloser{}}
	mockCalculator.EXPECT().Get(GBPDefaultURL).Return(r, nil)

	b, err := makeApiRequest(mockCalculator, GBPDefaultURL)

	if b != nil {
		t.Logf("want: nil, got: %v\n", b)
		t.Fail()
	}

	if !errors.Is(err, fake.SliceEmptyError) {
		t.Fatalf("want: %v, got: %v\n", fake.SliceEmptyError, err)
	}
}

const mid = 5.0

var sampleResponseBody = `{"table":"A","currency":"funt szterling","code":"GBP","rates":[{"no":"037/A/NBP/2023","effectiveDate":"2023-02-22","mid":` + strconv.FormatFloat(mid, 'f', -1, 64) + `}]}`
var invalidResponseBody = `{"table":"A","currency":"funt szterling","code":"GBP","rates":[{"no":"037/A/NBP/2023","effectiveDate":"2023-02-22","mid": "5.0"}]}` //string instead of float given

func TestGetCurrentRatePositive(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockCalculator := mock_calculator.NewMockHttpGetter(ctrl)
	body := ioutil.NopCloser(bytes.NewReader([]byte(sampleResponseBody)))
	r := &http.Response{StatusCode: 200, Body: body}

	mockCalculator.EXPECT().Get(GBPDefaultURL).Return(r, nil)

	v, err := GetCurrentRate(mockCalculator, "GBP", GBPDefaultURL)

	if v == 0 {
		t.Logf("want: %v, got: %v\n", mid, v)
		t.Fail()
	}

	if err != nil {
		t.Fatalf("want: nil, got: %v\n", err)
	}
}

func TestGetCurrentRateRequestError(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockCalculator := mock_calculator.NewMockHttpGetter(ctrl)
	r := &http.Response{StatusCode: 400, Body: http.NoBody}

	mockCalculator.EXPECT().Get(GBPDefaultURL).Return(r, nil)

	v, err := GetCurrentRate(mockCalculator, "GBP", GBPDefaultURL)

	if v != 0 {
		t.Logf("want: %v, got: %v\n", mid, v)
		t.Fail()
	}

	if err == nil {
		t.Fatalf("want: err, got: %v\n", err)
	}
}
func TestGetCurrentRateJsonUnmarshalError(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockCalculator := mock_calculator.NewMockHttpGetter(ctrl)
	body := ioutil.NopCloser(bytes.NewReader([]byte(invalidResponseBody)))
	r := &http.Response{StatusCode: 200, Body: body}

	mockCalculator.EXPECT().Get(GBPDefaultURL).Return(r, nil)

	v, err := GetCurrentRate(mockCalculator, "GBP", GBPDefaultURL)

	if v != 0 {
		t.Logf("want: %v, got: %v\n", 0, v)
		t.Fail()
	}

	if err == nil {
		t.Fatalf("want: err, got: %v\n", err)
	}
}
func TestGetCurrentRateIncorrectResponseError(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockCalculator := mock_calculator.NewMockHttpGetter(ctrl)
	body := ioutil.NopCloser(bytes.NewReader([]byte(sampleResponseBody)))
	r := &http.Response{StatusCode: 200, Body: body}

	mockCalculator.EXPECT().Get(GBPDefaultURL).Return(r, nil)

	v, err := GetCurrentRate(mockCalculator, "USD", GBPDefaultURL)

	if v != 0 {
		t.Logf("want: %v, got: %v\n", 0, v)
		t.Fail()
	}

	if !errors.Is(err, InvalidCurrencyReceivedErr) {
		t.Fatalf("want: %v, got: %v\n", InvalidCurrencyReceivedErr, err)
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
