package locate

import (
	"go-oss-demo/data/objects"
	"go-oss-demo/rabbitmq"
	"go-oss-demo/utils"
	"os"
	"strconv"
)

func Locate(name string) bool {
	_, e := os.Stat(name)
	return !os.IsNotExist(e)
}

func StartLocate() {
	mq := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer mq.Close()
	mq.Bind("dataServers")
	c := mq.Consume()
	for msg := range c {
		s, e := strconv.Unquote(string(msg.Body))
		if e != nil {
			panic(e)
		}
		if Locate(objects.Dir + s) {
			mq.Send(msg.ReplyTo, utils.GetLocalIP()+":"+os.Getenv("LISTEN_PORT"))
		}
	}
}
