package heartbeat

import (
	"go-oss-demo/rabbitmq"
	"os"
	"time"
)

func StartHeartbeat() {
	mq := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer mq.Close()
	for {
		mq.Publish("apiServers", os.Getenv("LISTEN_ADDR"))
		time.Sleep(time.Second * 5)
	}
}
