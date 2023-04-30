package main

import (
	"time"
	"web/service/api"
	"web/service/tcp"
)

func main() {
	go tcp.TCP_go()
	time.Sleep(time.Second*2)
	go api.GetApi()
	select {}
}
