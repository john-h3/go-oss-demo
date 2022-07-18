package locate

import (
	"encoding/json"
	"go-oss-demo/rabbitmq"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	url := Locate(strings.Split(r.URL.EscapedPath(), "/")[2])
	if len(url) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	bytes, _ := json.Marshal(url)
	w.Write(bytes)
}

func Locate(name string) string {
	mq := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	mq.Publish("dataServers", name)
	c := mq.Consume()
	go func() {
		time.Sleep(time.Second)
		mq.Close()
	}()
	msg := <-c
	url, _ := strconv.Unquote(string(msg.Body))
	return url
}
