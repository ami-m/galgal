package controllers

import (
	"dropit/externalApis/apis"
	"dropit/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResolveAddress Resolve address from text search and returns an address object, by applting
// to the GeoApiFy api
func ResolveAddress(c *gin.Context) {
	var b utils.Map
	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
	}

	json.Unmarshal(body, &b)
	geoApiFy := apis.GeoApiFy{
		Url:            "https://api.geoapify.com/v1/geocode/search",
		QueryString:    b["address"].(string),
		ResponseFormat: "json",
	}

	res, err := geoApiFy.SendRequest()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, res)
}
