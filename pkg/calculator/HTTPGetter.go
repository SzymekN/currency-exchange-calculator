package calculator

import "net/http"

type HttpGetter interface {
	Get(url string) (resp *http.Response, err error)
}

type DefaultHttpGetter struct {
}

func (d DefaultHttpGetter) Get(url string) (resp *http.Response, err error) {
	return http.Get(url)
}
