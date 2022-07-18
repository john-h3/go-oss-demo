package objectstream

import (
	"fmt"
	"io"
	"net/http"
)

type PutStream struct {
	writer *io.PipeWriter
	ch     chan error
}

func NewPutStream(server, object string) *PutStream {
	reader, writer := io.Pipe()
	ch := make(chan error)
	go func() {
		req, _ := http.NewRequest("PUT", "http://"+server+"/objects/"+object, reader)
		c := http.Client{}
		resp, e := c.Do(req)
		if e == nil && resp.StatusCode != http.StatusOK {
			e = fmt.Errorf("dataServer return http code %d", resp.StatusCode)
		}
		ch <- e
	}()
	return &PutStream{writer, ch}
}

func (w *PutStream) Write(p []byte) (n int, e error) {
	return w.writer.Write(p)
}

func (w *PutStream) Close() error {
	w.writer.Close()
	return <-w.ch
}

type GetStream struct {
	reader io.Reader
}

func newGetStream(url string) (*GetStream, error) {
	r, e := http.Get(url)
	if e != nil {
		return nil, e
	}
	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("dataServer return http code %d", r.StatusCode)
	}
	return &GetStream{r.Body}, nil
}

func NewGetStream(server, object string) (*GetStream, error) {
	if server == "" || object == "" {
		return nil, fmt.Errorf("invalid server %s object %s", server, object)
	}
	return newGetStream("http://" + server + "/objects/" + object)
}

func (r *GetStream) Read(p []byte) (n int, e error) {
	return r.reader.Read(p)
}
