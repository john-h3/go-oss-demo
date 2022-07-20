package heartbeat

import (
	"go-oss-demo/rabbitmq"
	"go-oss-demo/utils"
	"os"
	"time"
)

func StartHeartbeat() {
	mq := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer mq.Close()
	// if we don't consume, the queue will not be deleted by rabbitmq when the application down
	_ = mq.Consume()
	for {
		mq.Publish("apiServers", utils.GetLocalIP()+":"+os.Getenv("LISTEN_PORT"))
		time.Sleep(time.Second * 5)
	}
}
