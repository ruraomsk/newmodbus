package main

import (
	"time"

	"github.com/ruraomsk/newmodbus/client"
	"github.com/ruraomsk/newmodbus/server"
)

func main() {
	go server.Start()
	go client.Start()
	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
	}
}
