package main

import (
	"go-oss-demo/data/heartbeat"
	"go-oss-demo/data/locate"
	"go-oss-demo/data/objects"
	"log"
	"net/http"
	"os"
)

func main() {
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("LISTEN_PORT"), nil))
}
