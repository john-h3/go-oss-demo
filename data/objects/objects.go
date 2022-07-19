package objects

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var Dir = os.Getenv("STORAGE_DIR") + "/"

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
	f, e := os.Create(Dir + strings.Split(r.URL.EscapedPath(), "/")[2])
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError) // 500
		return
	}
	defer f.Close()
	io.Copy(f, r.Body)
}

func get(w http.ResponseWriter, r *http.Request) {
	f, e := os.Open(Dir + strings.Split(r.URL.EscapedPath(), "/")[2])
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}
	defer f.Close()
	io.Copy(w, f)
}
