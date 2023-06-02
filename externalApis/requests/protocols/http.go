package protocols

import (
	"dropit/externalApis/requests/formats"
	"dropit/utils"
	"fmt"
	"net/http"
	"time"
)

type Http struct {
	Common
	Method string
	Url    string
}

func (h Http) GenerateRequest(format string) (interface{}, error) {
	var err error
	req, err := http.NewRequest(h.Method, h.Url, nil)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Couldn't create new request, %s", err))
	}

	f := formats.Factory(format, utils.Map{"body": h.Body.(utils.Map)})
	f.Parse(req)
	client := http.Client{Timeout: 30 * time.Second}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Request was failed, %s", err))
	}

	return res, nil
}
