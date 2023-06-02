package globals

import (
	"dropit/files"
	"dropit/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var globals utils.Map

type Holiday struct {
	Name    string `json:"name"`
	Date    string `json:"date"`
	Country string `json:"country"`
}

type Holidays struct {
	Holidays []Holiday `json:"holidays"`
}

const (
	HOLIDAYS = "holidays"
)

// InitGlobals Assigns all system required global variables. Called by app Init() function
func InitGlobals() {
	globals = utils.Map{
		HOLIDAYS: extrctHolidays(),
	}
}

// extrctHolidays This function reads the holiday.json and extract it into map where the
// date+country is the key and the value is the holiday struct itself
func extrctHolidays() map[string]Holiday {
	filePath, err := files.GetFilePath(files.HOLIDAYS)
	if err != nil {
		panic(err)
	}

	byteValue, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var holidays Holidays
	mappedHolidays := map[string]Holiday{}
	json.Unmarshal(byteValue, &holidays)
	for _, holiday := range holidays.Holidays {
		mappedHolidays[holiday.Date] = holiday
	}

	return mappedHolidays
}

// GetGlobal Returns the global variable value by key name, otherwise if key not exists returns error
func GetGlobal(name string) (interface{}, error) {
	if _, ok := globals[name]; !ok {
		return "", fmt.Errorf(fmt.Sprintf("No global with name %s", name))
	}

	return globals[name], nil
}
