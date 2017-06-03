package connections

import (
	log "github.com/Sirupsen/logrus"
	"github.com/garyburd/redigo/redis"
	"github.com/liam-lai/ptt-alertor/myutil"
)

func Redis() redis.Conn {
	config := myutil.Config("redis")
	conn, err := redis.Dial("tcp", config["host"]+":"+config["port"])
	if err != nil {
		log.Fatal(err)
	}
	return conn
}