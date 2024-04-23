package main

import (
	"time"
	"web/service/api"
	"web/service/tcp"
)

func main() {
	port := 8081
	go tcp.TCPGo(port+1)
	time.Sleep(time.Second*2)
	go api.GetApi(port)
	select {}
}
