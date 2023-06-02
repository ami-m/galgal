package connections

import (
	"dropit/configs"
	"fmt"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mySql struct {
	common
}

func (m mySql) Connect() interface{} {
	if m.client == nil {
		m.mu.Lock()
		if m.client == nil {
			conf := configs.GetConfig()
			connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", conf.MysqlUser, conf.MysqlPassword, conf.MysqlHost, conf.MysqlPort, conf.MysqlDB)
			fmt.Println("Connection String: " + connectionString)
			log.Info("Connecting to mysql with the following details: ", connectionString)
			var err error
			m.client, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
			if err != nil {
				m.mu.Unlock()
				panic(fmt.Sprintf("Couldn't connect to DB, due to: %s", err.Error()))
			}
			m.mu.Unlock()
		}
	}

	return m.client
}
