package fake

import (
	"testing"
)

func TestReadCloser_Read(t *testing.T) {
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		rc      ReadCloser
		args    args
		wantN   int
		wantErr bool
	}{
		{
			name:    "Positive case",
			rc:      ReadCloser{},
			args:    args{p: []byte{15}},
			wantN:   1,
			wantErr: false,
		},
		{
			name:    "Slice empty test case",
			rc:      ReadCloser{},
			args:    args{p: []byte{}},
			wantN:   0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, err := tt.rc.Read(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadCloser.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("ReadCloser.Read() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestReadCloser_Close(t *testing.T) {
	tests := []struct {
		name    string
		rc      ReadCloser
		wantErr bool
	}{
		{name: "Positive case", rc: ReadCloser{}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.rc.Close(); (err != nil) != tt.wantErr {
				t.Errorf("ReadCloser.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
