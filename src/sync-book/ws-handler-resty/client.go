package handler_resty

import "github.com/go-resty/resty/v2"

var client = resty.New()

func Get(uri string) (*resty.Response, error) {
	return client.R().
		EnableTrace().
		Get(uri)
}
