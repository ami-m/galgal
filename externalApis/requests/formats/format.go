package formats

import (
	"dropit/utils"
	"net/http"
)

type Format interface {
	//Parse Parses the request body by the desired format
	Parse(request *http.Request) interface{}
}

type Common struct {
	Body utils.Map
}

const (
	URL_QUERY = "urlQuery"
)

func Factory(name string, props utils.Map) Format {
	switch name {
	case URL_QUERY:
		return urlQuery{
			Common{
				Body: props["body"].(utils.Map),
			},
		}
	default:
		panic("Wrong request format")
	}
}
