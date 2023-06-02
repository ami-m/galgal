package connections

import "sync"

const MYSQL = "mysql"

var drivers map[string]Connection = map[string]Connection{
	MYSQL: mySql{common{mu: &sync.Mutex{}}},
}

type Connection interface {
	Connect() interface{}
}

type common struct {
	client interface{}
	mu     *sync.Mutex
}

func GetDriver(name string) interface{} {
	switch name {
	case MYSQL:
		return drivers[name].Connect()
	default:
		panic("Wrong driver name was provided")
	}
}
