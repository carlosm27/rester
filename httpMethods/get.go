package httpMethods

import (
	"log"

	"github.com/go-resty/resty/v2"
)

func GetUri(uri string) (resp *resty.Response, err error) {

	client := resty.New()

	resp, err = client.R().
		EnableTrace().
		Get(uri)
	if err != nil {
		log.Println(err)
	}

	return
}
