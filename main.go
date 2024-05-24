package main

import (
	"time"

	"github.com/ruraomsk/newmodbus/client"
	"github.com/ruraomsk/newmodbus/server"
)

func main() {
	server.Start()
	client.Start()
	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
	}
}
