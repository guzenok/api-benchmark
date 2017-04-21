// http_test.go
package serv

import (
	"bytes"
	"testing"
	"time"
)

// Test data

func BenchmarkHttp(b *testing.B) {
	for serverName, handleCreate := range HttpHandlerList {
		for decoderName, decodeFunc := range DecodesList {
			encodeFunc, _ := EncodesList[decoderName]
			// запуск http-сервера
			stop := HttpServe(URL, handleCreate(decodeFunc, encodeFunc))
			time.Sleep(time.Second) // пусть устаканится
			// запуск теста
			b.Run("API("+serverName+"+"+decoderName+")", func(b *testing.B) {
				data := NewTransferRequest().ToJSON()
				for i := 0; i < b.N; i++ {
					body := bytes.NewReader(data)
					HttpGet(URL, body)
				}
			})
			// остановка http-сервера
			stop()
			time.Sleep(time.Second) // пусть устаканится
		}
	}
}
