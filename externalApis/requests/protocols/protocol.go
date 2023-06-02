package protocols

import (
	"dropit/utils"
)

type Protocol interface {
	// Return a request object
	GenerateRequest(format string) (interface{}, error)
}

type Common struct {
	Body interface{}
}

const (
	HTTP = "http"
)

// Returns Protocol interface object
func Factory(name string, props utils.Map) Protocol {
	switch name {
	case HTTP:
		return Http{
			Common: Common{Body: props["body"].(utils.Map)},
			Method: props["method"].(string),
			Url:    props["url"].(string),
		}
	default:
		panic("Wrong protocol name")
	}
}
