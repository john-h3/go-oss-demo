package main

import (
	"go-oss-demo/api/heartbeat"
	"go-oss-demo/api/locate"
	"go-oss-demo/api/objects"
	"log"
	"net/http"
	"os"
)

func main() {
	go heartbeat.ListenHeartBeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("LISTEN_PORT"), nil))
}
