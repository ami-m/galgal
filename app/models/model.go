package models

import (
	"dropit/databases/connections"

	"gorm.io/gorm"
)

// Model Prime interface for all models. All other model actions interfaces inherit it.
type Model interface {
	GetTableName() string
}

// Retrieval Interface to models who need to fetch records
type Retrieval interface {
	Model
	// GetRecords Function, where models fetch records from DB
	GetRecords(args ...interface{}) ([]Model, error)
}

// Insertion Interface to models who need to insert records
type Insertion interface {
	Model
	// GetRecords Function, where models insert records to DB
	SetRecord(args ...interface{}) ([]Model, error)
}

// Deletion Interface to models who need to delete records
type Deletion interface {
	Model
	// GetRecords Function, where models delete records from DB
	DeleteRecord(id int64) error
}

// GetDriver Returns the MYSQL driver to all models
func GetDriver() *gorm.DB {
	return connections.GetDriver(connections.MYSQL).(*gorm.DB)
}

// ConvertResultToModelArray convert the obtained results from query to array of Model interface
// In order to enforce the abstraction in models' return collection
func ConvertResultToModelArray[T Model](result []T) []Model {
	var models []Model
	for _, entity := range result {
		models = append(models, entity)
	}
	return models
}
