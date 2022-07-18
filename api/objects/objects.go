package objects

import (
	"fmt"
	"go-oss-demo/api/heartbeat"
	"go-oss-demo/api/locate"
	"go-oss-demo/api/objectstream"
	"io"
	"log"
	"net/http"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodPut {
		put(w, r)
		return
	}
	if m == http.MethodGet {
		get(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func put(w http.ResponseWriter, r *http.Request) {
	object := strings.Split(r.URL.EscapedPath(), "/")[2]
	c, e := storeObject(r.Body, object)
	if e != nil {
		log.Println(e)
	}
	w.WriteHeader(c)
}

func storeObject(r io.Reader, object string) (int, error) {
	stream, e := putStream(object)
	if e != nil {
		return http.StatusServiceUnavailable, e
	}
	io.Copy(stream, r)
	e = stream.Close()
	if e != nil {
		return http.StatusInternalServerError, e
	}
	return http.StatusOK, nil
}

func putStream(object string) (*objectstream.PutStream, error) {
	server := heartbeat.ChooseRandomDataServer()
	if server == "" {
		return nil, fmt.Errorf("no dataServer available")
	}
	return objectstream.NewPutStream(server, object), nil
}

func get(w http.ResponseWriter, r *http.Request) {
	object := strings.Split(r.URL.EscapedPath(), "/")[2]
	stream, e := getStream(object)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	io.Copy(w, stream)
}

func getStream(object string) (*objectstream.GetStream, error) {
	server := locate.Locate(object)
	if server == "" {
		return nil, fmt.Errorf("object %s locate fail", object)
	}
	return objectstream.NewGetStream(server, object)
}
