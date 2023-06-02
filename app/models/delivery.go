package models

import (
	"dropit/app/globals"
	"dropit/utils"
	"fmt"
	"strings"
	"sync"
	"time"

	"gorm.io/gorm"
)

const (
	WEEKLY                    = "weekly"
	DAILY                     = "daily"
	DELIVER_BUSINESS_CAPACITY = 10
)

type Delivery struct {
	ID         int64
	Status     string `gorm:"type:enum('NEW','ARRIVED','CANCELED');default:'NEW'"`
	TimeSlotID int64
	TimeSlot   TimeSlot
	CreatedAt  time.Time
	UpdateAt   *time.Time
	DeletedAt  gorm.DeletedAt
}

/*
  - This map responsible to store mutex by timeSlotId and therefore letting other request with
    different timeSlotId to keep their execution
*/
var timeSlotMutex map[string]*sync.Mutex = map[string]*sync.Mutex{}

/*
  - This mutex responsible to allow only one request at a time to ask the amount of the daily deliveries
    and by doing it, preventing from other request to insert more deliveries ubiquitously while the locking request
    still running *
*/
var deliveryCapitcityMutex *sync.Mutex = &sync.Mutex{}

func (d Delivery) GetRecords(args ...interface{}) ([]Model, error) {
	switch args[0].(string) {
	case DAILY:
		return GetDailyDeliveries()
	case WEEKLY:
		return GetWeeklyDeliveries()
	default:
		return []Model{}, nil
	}
}

func (d Delivery) SetRecord(args ...interface{}) (Model, error) {
	var timeSlot TimeSlot
	// Fetch request's timeSlotId param
	timeSlotId := args[0].(int64)

	/** Here we are checking if the incoming timeSlotId is referenced inside the
	timeSlotMutex array */
	if _, ok := timeSlotMutex[string(timeSlotId)]; !ok {
		timeSlotMutex[string(timeSlotId)] = &sync.Mutex{}
	}

	/** Now we are lock the code only for requests with the specific timeSlotId arrived from the request
	,hence, letting other requests to keep moving on throughout the code freely */
	timeSlotMutex[string(timeSlotId)].Lock()
	defer timeSlotMutex[string(timeSlotId)].Unlock()
	res := GetDriver().First(&timeSlot, timeSlotId)
	if res.Error != nil {
		return nil, res.Error
	}

	if err := validaeTimeSlot(timeSlot); err != nil {
		return nil, err
	}

	if err := validateDeliveriesBuisinessCapicity(); err != nil {
		return nil, err
	}

	d.TimeSlotID = timeSlotId
	/** We are calling here to a new transaction commit so if one of the actions:
	    1. Creating new delivery
		2. Updating time slots num_of_deliveries
		will fail the all transaction will be rolled back */
	err := GetDriver().Transaction(func(tx *gorm.DB) error {
		if err := GetDriver().Create(&d).Error; err != nil {
			return err
		}

		if res := GetDriver().Model(&timeSlot).Update("num_of_deliveries", timeSlot.NumOfDeliveries+1); res.Error != nil {
			return res.Error
		}

		return nil
	})

	return d, err
}

func (d Delivery) DeleteRecord(id int64) error {
	d.ID = id
	err := GetDriver().Transaction(func(tx *gorm.DB) error {
		GetDriver().Model(&d).Update("status", "CANCELED")
		if err := GetDriver().Model(&d).Update("status", "CANCELED").Error; err != nil {
			return err
		}

		if err := GetDriver().Delete(&d, id).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}

// GetDailyDeliveries Retrieves all the daily deliveries
func GetDailyDeliveries() ([]Model, error) {
	var deliveries []Delivery
	res := GetDriver().Where("created_at BETWEEN ? AND ?",
		fmt.Sprintf("%s 00:00:00", time.Now().Format(utils.DateFormat)),
		fmt.Sprintf("%s 23:59:59", time.Now().Format(utils.DateFormat)),
	).Find(&deliveries)

	if res.Error != nil {
		return nil, res.Error
	}

	return ConvertResultToModelArray(deliveries), nil
}

// GetWeeklyDeliveries Retrieves all the weekly deliveries
func GetWeeklyDeliveries() ([]Model, error) {
	var deliveries []Delivery
	res := GetDriver().Where("created_at BETWEEN ? AND ?",
		fmt.Sprintf("%s 00:00:00", time.Now().AddDate(0, 0, -7).Format(utils.DateFormat)),
		fmt.Sprintf("%s 23:59:59", time.Now().Format(utils.DateFormat)),
	).Find(&deliveries)

	if res.Error != nil {
		return nil, res.Error
	}

	return ConvertResultToModelArray(deliveries), nil
}

func (d Delivery) GetTableName() string {
	return "deliveries"
}

// validateDeliveriesBuisinessCapicity Validate that deliveries business capacity (10) is between boundaries
func validateDeliveriesBuisinessCapicity() error {
	// Here We are locking the code with mutex, due o the fact that several requests can change the capacity of deliveries simultaneously
	deliveryCapitcityMutex.Lock()
	defer deliveryCapitcityMutex.Unlock()
	dayilyDeliveries, _ := GetDailyDeliveries()
	if len(dayilyDeliveries) >= DELIVER_BUSINESS_CAPACITY {
		return fmt.Errorf("Deliveries Business capacity (10) exceeded")
	}

	return nil
}

// validaeTimeSlot Validates whether time slots vaccant (attached to at most 2 deliveries) and its date is not overlapping with one of the holidays
func validaeTimeSlot(t TimeSlot) error {
	holidayName, isHoliday := isTimeSlotDateIsHoliday(t)
	if isHoliday {
		return fmt.Errorf(fmt.Sprintf("TimeSlot id %d is overlapping with holiday %s", t.ID, holidayName))
	}

	if !isTimeSlotVacant(t) {
		return fmt.Errorf(fmt.Sprintf("TimeSlot id %d has already appointed to 2 other deliveries", t.ID))
	}

	return nil
}

// isTimeSlotVacant Checks whether time slot is not attached to 2 different deliveries
func isTimeSlotVacant(t TimeSlot) bool {
	if t.NumOfDeliveries >= MAX_NUM_OF_DELIERIES {
		return false
	}

	return true
}

// isTimeSlotDateIsHoliday Checks whether time slot date is not overlapping with one of the holidays dates
func isTimeSlotDateIsHoliday(t TimeSlot) (string, bool) {
	/** The Date format is yyyy-mm-dd h:i:s, but the holidays global map keys ate only dates
	, hence we splitting here the date and taking only the first part */
	deliverDate := strings.Split(string(t.StartTime), " ")[0]
	holidays, _ := globals.GetGlobal(globals.HOLIDAYS)
	if _, ok := holidays.(map[string]globals.Holiday)[deliverDate]; !ok {
		return "", false
	}

	return holidays.(map[string]globals.Holiday)[deliverDate].Name, true
}
