package formats

import "net/http"

type urlQuery struct {
	Common
}

func (u urlQuery) Parse(request *http.Request) interface{} {
	q := request.URL.Query()
	for key, param := range u.Body {
		q.Add(key, param.(string))
	}

	request.URL.RawQuery = q.Encode()
	return request.URL.String()
}
