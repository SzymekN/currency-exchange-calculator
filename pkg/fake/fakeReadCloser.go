package fake

import (
	"errors"
)

var SliceEmptyError = errors.New("Slice empty error")

type ReadCloser struct {
}

func (rc ReadCloser) Read(p []byte) (n int, err error) {
	if len(p) == 0 || len(p) == 512 {
		return 0, SliceEmptyError
	}
	return len(p), nil
}

func (rc ReadCloser) Close() error {
	return nil
}
