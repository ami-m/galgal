package formats

import (
	"dropit/app/models"
	"dropit/utils"
	"net/http"
)

type Format interface {
	//Parse Parses the request body by the desired format
	Parse(response http.Response) (interface{}, error)
}

type Common struct {
	Body utils.Map
}

const (
	JSON = "json"
)

// Factory Returns response Format object
func Factory(name string, prop utils.Map) Format {
	switch name {
	case JSON:
		return Json{Model: prop["model"].(models.Model)}
	default:
		panic("Wrong Response format")
	}
}
