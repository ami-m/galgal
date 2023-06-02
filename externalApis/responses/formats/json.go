package formats

import (
	"dropit/app/models"
	"encoding/json"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Json struct {
	Model models.Model
}

func (j Json) Parse(response http.Response) (interface{}, error) {
	defer response.Body.Close()
	// read body
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	log.Infof("response body: %s", string(responseBody))
	res, _ := json.Marshal(string(responseBody))
	json.Unmarshal(res, &j.Model)
	return j.Model, nil
}
