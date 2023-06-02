package configs

import (
	"dropit/utils"
	"sync"

	log "github.com/sirupsen/logrus"
)

// Config holds all configurable fields
type Config struct {
	LogLevel       log.Level // log messages from this level and lower (lower = more important)
	Protocol       string
	Host           string
	Port           string
	MysqlHost      string
	MysqlPort      string
	MysqlUser      string
	MysqlPassword  string
	MysqlDB        string
	GeoApiFyApiKey string
	GeoApiFyUrl    string
}

var conf *Config
var mu *sync.Mutex = &sync.Mutex{}

func GetConfig() *Config {
	// In order to acquire config Singelton instance we will user double locking in case
	// the conf variable were initialized before the predecessor gouroutine acquire the mutex
	if conf == nil {
		mu.Lock()
		if conf == nil {
			conf = NewConfig()
		}
		mu.Unlock()
	}

	return conf
}

// NewConfig builds the Config instance from os environment variables
func NewConfig() *Config {
	logLevelName := utils.GetEnv("LOG_LEVEL", "info")
	logLevelId, err := log.ParseLevel(logLevelName)
	if err != nil {
		logLevelId = log.InfoLevel
	}

	// intervalDuration, _ := strconv.Atoi(utils.GetEnv("INTERVAL_DURATION", "1"))
	return &Config{
		LogLevel:       logLevelId,
		Protocol:       utils.GetEnv("PROTOCOL", "http"),
		Host:           utils.GetEnv("HOST", "127.0.0.1"),
		Port:           utils.GetEnv("PORT", "8000"),
		MysqlHost:      utils.GetEnv("MYSQL_HOST", "127.0.0.1"),
		MysqlPort:      utils.GetEnv("MYSQL_PORT", "3306"),
		MysqlUser:      utils.GetEnv("MYSQL_USER", "root"),
		MysqlPassword:  utils.GetEnv("MYSQL_PASSWORD", ""),
		MysqlDB:        utils.GetEnv("MYSQL_DB", "dropit"),
		GeoApiFyUrl:    utils.GetEnv("GEO_API_FY_URL", "https://api.geoapify.com/v1/geocode"),
		GeoApiFyApiKey: utils.GetEnv("GEO_API_FY_KEY", ""),
	}
}
