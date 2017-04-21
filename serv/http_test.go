// http_test.go
package serv

import (
	"bytes"
	"testing"
	"time"
)

// Test data
var (
	URL = ":3000"

	testServers = map[string]HandleCreator{
		"Gin": GinHandler,
		"Chi": ChiHandler,
	}

	testDecodes = map[string]DecodeFunc{
		"Standart": DecodeStandart,
		"Buger":    DecodeBuger,
	}

	testEncodes = map[string]EncodeFunc{
		"Standart": EncodeStandart,
		"Buger":    EncodeBuger,
	}
)

func BenchmarkHttp(b *testing.B) {
	for serverName, handleCreate := range testServers {
		for decoderName, decodeFunc := range testDecodes {
			encodeFunc, _ := testEncodes[decoderName]
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
