package seeds

import (
	"dropit/app/models"
	"dropit/files"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type TimeSlot struct {
	Common
}

func (t TimeSlot) Run() {
	var courierAvailableTimeslots models.CourierAvailableTimeslots
	filePath, err := files.GetFilePath(files.COURIER_AVAILABLE_TIMESLOTS)
	if err != nil {
		panic(err)
	}

	byteValue, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.LogError(
			fmt.Errorf(
				fmt.Sprintf("Couldn't read file %s, %s", files.COURIER_AVAILABLE_TIMESLOTS, err.Error()),
			),
		)
	}

	json.Unmarshal(byteValue, &courierAvailableTimeslots)
	result := models.
		GetDriver().
		Create(courierAvailableTimeslots.TimeSlots)
	if result.Error != nil {
		t.LogError(result.Error)
		return
	}

	t.LogSuccess(int(result.RowsAffected))
}
