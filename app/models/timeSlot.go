package models

import (
	"dropit/utils"
	"fmt"
	"time"

	"gorm.io/datatypes"
)

const (
	MAX_NUM_OF_DELIERIES = 2
)

type TimeSlot struct {
	ID              int64
	StartTime       utils.JsonTime `gorm:"not null" json:"start_time"`
	EndTime         utils.JsonTime `gorm:"not null" json:"end_time"`
	PostCodes       datatypes.JSON `gorm:"type:json" json:"supported_postcodes"`
	NumOfDeliveries int8           `gorm:"default:0" json:"num_of_deliveries"`
	CreatedAt       *time.Time     `json:"created_at"`
	UpdateAt        *time.Time     `json:"updated_at"`
}

// This struct primarily purpose is to unmarshal the time_slots.json file
type CourierAvailableTimeslots struct {
	TimeSlots []TimeSlot `json:"courier_available_timeslots"`
}

func (t TimeSlot) GetRecords(args ...interface{}) ([]Model, error) {
	var timeSlots []TimeSlot
	res := GetDriver().Raw(
		fmt.Sprintf("SELECT * FROM %s WHERE JSON_CONTAINS(post_codes, ?, '$') AND nun_of_deliveries < 2", t.GetTableName()),
		"\""+args[0].(Address).PostCode+"\"",
	).Find(&timeSlots)

	if res.Error != nil {
		return nil, res.Error
	}

	return ConvertResultToModelArray(timeSlots), nil
}

func (t TimeSlot) GetTableName() string {
	return "time_slots"
}
