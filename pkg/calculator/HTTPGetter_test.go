package calculator

import (
	"net/http"
	"testing"
)

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
			args:     args{url: GBPDefaultURL},
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
