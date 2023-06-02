package files

import (
	"dropit/utils"
	"fmt"
	"path/filepath"
)

const (
	COURIER_AVAILABLE_TIMESLOTS = "courier_available_timeslots"
	HOLIDAYS                    = "holidays"
)

// filesPaths Map to store all system files paths
var filesPaths = map[string]string{
	"courier_available_timeslots": utils.StoragePath(filepath.Join("app", COURIER_AVAILABLE_TIMESLOTS+".json")),
	"holidays":                    utils.StoragePath(filepath.Join("app", HOLIDAYS+".json")),
}

// GetFilePath Returns file path out of the files index
func GetFilePath(fileName string) (string, error) {
	if _, ok := filesPaths[fileName]; !ok {
		return "", fmt.Errorf(fmt.Sprintf("The file %s is not exits", fileName))
	}

	return filesPaths[fileName], nil
}
