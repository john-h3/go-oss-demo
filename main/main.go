package main

import (
	"go-oss-demo/main/objects"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := os.Getenv("LISTEN_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	log.Println("listening on", addr)
	dir := os.Getenv("STORAGE_DIR")
	if dir == "" {
		dir = objects.Dir
	} else {
		objects.Dir = dir
	}
	log.Println("storage dir:", dir)
	_, e := os.Stat(dir)
	if e != nil {
		if os.IsNotExist(e) {
			_ = os.MkdirAll(objects.Dir, os.ModeDir|os.ModePerm)
		} else {
			log.Fatal(e)
		}
	}
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(addr, nil))
}
