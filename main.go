// main.go
package main

import (
	"bytes"
	"flag"
	"log"
	"time"

	"github.com/guzenok/api-benchmark/serv"
)

var (
	command = flag.String("act", "all", "")

	URL = ":3000"

	testServers = map[string]serv.HandleCreator{
		"Gin": serv.GinHandler,
		"Chi": serv.ChiHandler,
	}

	testDecodes = map[string]serv.DecodeFunc{
		"Standart": serv.DecodeStandart,
		"Buger":    serv.DecodeBuger,
	}

	testEncodes = map[string]serv.EncodeFunc{
		"Standart": serv.EncodeStandart,
		"Buger":    serv.EncodeBuger,
	}
)

func main() {
	flag.Parse()

	TestAll(1000000)

}

func TestAll(N int) {
	for serverName, handleCreate := range testServers {
		for decoderName, decodeFunc := range testDecodes {
			encodeFunc, _ := testEncodes[decoderName]
			log.Println("==================================================")
			log.Printf("Testing %s with %s\n", serverName, decoderName)
			// запуск сервера
			stop := serv.HttpServe(URL, handleCreate(decodeFunc, encodeFunc))
			time.Sleep(time.Second) // пусть устаканится
			// время начала
			start := time.Now()
			// тело запроса
			for i := 0; i < N; i++ {
				b := serv.NewTransferRequest().ToJSON()
				body := bytes.NewReader(b)
				serv.HttpGet(URL, body)
			}
			// время завершения
			finish := time.Now()
			// остановка сервера
			stop()
			time.Sleep(time.Second) // пусть устаканится
			log.Printf(" takes %d microSecond\n", int64(finish.Sub(start)/time.Microsecond)/int64(N))
		}
	}
}
