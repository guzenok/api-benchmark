// handler.go
package serv

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pressly/chi"
	"gopkg.in/gin-gonic/gin.v1"
)

type HandleCreator func(DecodeFunc, EncodeFunc) http.Handler

func ChiHandler(decode DecodeFunc, encode EncodeFunc) http.Handler {
	r := chi.NewRouter()
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {

		resp, err := convert(r.Body, decode, encode)
		if err != nil {
			// c.AbortWithStatus(http.DefaultTransport)
		}
		w.Write(resp)
	})
	return r
}

func GinHandler(decode DecodeFunc, encode EncodeFunc) http.Handler {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.POST("/", func(c *gin.Context) {

		resp, err := convert(c.Request.Body, decode, encode)
		if err != nil {
			// c.AbortWithStatus(http.DefaultTransport)
		}
		c.Writer.Write(resp)
	})
	return r
}

func convert(body io.Reader, decode DecodeFunc, encode EncodeFunc) (resp []byte, err error) {
	resp = []byte("{}")
	buff, err := ioutil.ReadAll(body)
	if err != nil {
		log.Printf("ReqRead err: %s", err)
		return nil, err
	}
	// log.Printf("ReqRead: %s", buff)
	tr, err := decode(buff)
	if err != nil {
		log.Printf("Decode err: %s", err)
		return nil, err
	}
	resp, err = encode(tr)
	return resp, err
}
