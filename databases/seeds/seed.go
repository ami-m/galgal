package seeds

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type Seed interface {
	//Function to run the specific table seed
	Run()
}

type Common struct {
	Name string
}

// LogError Writes the seed termination error to log
func (c Common) LogError(err error) {
	log.Error(fmt.Sprintf("Error in %s seeds, %s", c.Name, err.Error()))
}

// LogSuccess Writes to log successful message and the number of records assigned to the DB table
func (c Common) LogSuccess(rowsAffected int) {
	log.Info(fmt.Sprintf("%s has finished successfully. Number of records created:  %d", c.Name, rowsAffected))
}
