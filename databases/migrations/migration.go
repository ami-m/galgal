package migrations

import (
	"dropit/app/models"
	"dropit/databases/connections"

	"gorm.io/gorm"
)

var driver interface{}

const (
	CREATE  = "create"
	DESTROY = "destroy"
)

func Run(action string) {
	driver = connections.GetDriver(connections.MYSQL)
	switch action {
	case CREATE:
		create()
	case DESTROY:
		destroy()
	default:
		panic("Wrong argument has been sent")
	}
}

// This function responsible for building the DB tables
func create() {
	driver.(*gorm.DB).AutoMigrate(&models.TimeSlot{}, &models.Delivery{})
}

// This function responsible to destroy all DB tables
func destroy() {
	driver.(*gorm.DB).Migrator().DropTable(&models.Delivery{}, models.TimeSlot{})
}
