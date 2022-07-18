package heartbeat

import (
	"go-oss-demo/rabbitmq"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

var dataServers = make(map[string]time.Time)
var mutex sync.Mutex

func ListenHeartBeat() {
	mq := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer mq.Close()
	mq.Bind("apiServers")
	c := mq.Consume()
	go removeExpiredDataServer()
	for msg := range c {
		dataServer, e := strconv.Unquote(string(msg.Body))
		if e != nil {
			panic(e)
		}
		mutex.Lock()
		dataServers[dataServer] = time.Now()
		mutex.Unlock()
	}
}

func removeExpiredDataServer() {
	for {
		time.Sleep(time.Second * 5)
		mutex.Lock()
		for dataServer, lastSeen := range dataServers {
			if time.Now().Sub(lastSeen) > time.Second*10 {
				delete(dataServers, dataServer)
			}
		}
		mutex.Unlock()
	}
}

func GetDataServers() []string {
	mutex.Lock()
	defer mutex.Unlock()
	ds := make([]string, 0)
	for dataServer, _ := range dataServers {
		ds = append(ds, dataServer)
	}
	return ds
}

func ChooseRandomDataServer() string {
	ds := GetDataServers()
	if len(ds) == 0 {
		return ""
	}
	return ds[rand.Intn(len(ds))]
}
