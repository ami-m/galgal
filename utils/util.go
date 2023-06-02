package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// InitEnv Extract the the .env file into global array to hold all the environment variables
func InitEnv(content string) {
	records := strings.Split(content, "\n")
	for _, record := range records[:len(records)-1] {
		s := strings.Split(record, "=")
		err := os.Setenv(s[0], s[1])
		if err != nil {
			panic(fmt.Sprintf("Fail to set env record, due to: %s", err))
		}
	}
}

// GetEnv Function to fetch the required parameter from the environment variables and return it back
func GetEnv(name string, defVal string) string {
	val := os.Getenv(name)
	if val == "" {
		return defVal
	}
	return val
}

// "yyyy-mm-dd h:i:s" format for json date time fields
var TimeFormat = "2006-01-02 15:04:05"
var DateFormat = "2006-01-02"

type JsonTime string

// MarshalJSON This function is called when json.Marshal encoder is executed on a JsonTime property.
// It converts the date string to "yyyy-mm-dd h:i:s" format
func (t JsonTime) MarshalJSON() ([]byte, error) {
	if len(t) == 0 {
		return json.Marshal(nil)
	}
	timeObj, _ := time.Parse(TimeFormat, string(t))
	esTime := timeObj.Format(DateFormat)
	return json.Marshal(esTime)
}

// This function responsible to return the full storage's file path
func StoragePath(relativePath string) string {
	storageDir := GetEnv("STORAGE_DIR", "storage")
	cwd, _ := os.Getwd()
	return filepath.Join(cwd, string(os.PathSeparator), storageDir, string(os.PathSeparator), relativePath)
}

type Map = map[string]interface{}
