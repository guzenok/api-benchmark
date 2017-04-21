// main.go
package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/guzenok/api-benchmark/serv"
	"github.com/pkg/profile"
)

var (
	pprof    = flag.String("pprof", "Mem", "Which profile use (CPU, Mem, Block, Trace or Mutex)")
	httpserv = flag.String("httpserv", "Gin", "Which type of http-server use (Gin or Chi)")
	parser   = flag.String("parser", "Buger", "Which type of http-server use (Standart or Buger)")
)

func main() {
	// разбор аргументов
	flag.Parse()

	// выбор профилировщика
	pprofOpt, found := map[string]func(*profile.Profile){
		"CPU":   profile.CPUProfile,
		"Mem":   profile.MemProfile,
		"Block": profile.BlockProfile,
		"Mutex": profile.MutexProfile,
		"Trace": profile.TraceProfile,
	}[*pprof]
	if !found {
		flag.PrintDefaults()
		return
	}
	defer profile.Start(profile.ProfilePath("."), pprofOpt).Stop()

	// выбор http-обработчика
	handleCreate, found := serv.HttpHandlerList[*httpserv]
	if !found {
		flag.PrintDefaults()
		return
	}

	// выбор json-парсера
	decodeFunc, found := serv.DecodesList[*parser]
	if !found {
		flag.PrintDefaults()
		return
	}
	encodeFunc := serv.EncodesList[*parser]

	// info
	fmt.Printf("Start %s at http://localhost%s (parser %s)\n", *httpserv, serv.URL, *parser)

	// OS Signal
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// start http-server
	stop := serv.HttpServe(serv.URL, handleCreate(decodeFunc, encodeFunc))
	defer stop()
	time.Sleep(time.Second) // пусть устаканится
	// бесконечные запросы к серверу
	data := serv.NewTransferRequest().ToJSON()
Infinity:
	for {
		select {
		case <-sigs:
			break Infinity
		default:
			break
		}
		body := bytes.NewReader(data)
		serv.HttpGet(serv.URL, body)
		time.Sleep(time.Millisecond)
	}
}
