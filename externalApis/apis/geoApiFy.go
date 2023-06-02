package apis

import (
	"dropit/configs"
	"dropit/externalApis/requests/formats"
	"dropit/externalApis/requests/protocols"
	"dropit/utils"
	"net/http"
	"net/url"
)

type GeoApiFy struct {
	QueryString    string
	Url            string
	ResponseFormat string
}

func (g GeoApiFy) SendRequest() (interface{}, error) {
	url, err := url.JoinPath(g.Url, "search")
	if err != nil {
		return nil, err
	}

	h := protocols.Factory(
		protocols.HTTP,
		utils.Map{
			"url":    url,
			"method": http.MethodGet,
			"body": utils.Map{
				"text":   g.QueryString,
				"forma":  g.ResponseFormat,
				"apiKey": configs.GetConfig().GeoApiFyApiKey,
			},
		},
	)

	res, err := h.GenerateRequest(formats.URL_QUERY)
	return res, err
}
