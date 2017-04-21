// http.go
package serv

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	URL = ":3000"
)

func HttpServe(addr string, handler http.Handler) func() {

	srv := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			// log.Printf("http server failed: %v", err)
		}
	}()

	return func() {
		srv.Close()
	}

}

func HttpGet(addr string, body io.Reader) {
	resp, err := http.Post("http://localhost"+addr, "application/json", body)
	if err != nil {
		log.Printf("POST ERR: %s\n", err)
		return
	}
	buff, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("RespRead err: %+v (%s)", err, buff)
	} else {
		// log.Printf("RespRead: %s", buff)
	}
	resp.Body.Close()
}
