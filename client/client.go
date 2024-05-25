package client

import (
	"fmt"
	"time"

	"github.com/ruraomsk/newmodbus/modbus"
)

var client *modbus.ModbusClient
var err error

func Start() {
	fmt.Println("Client start")
	client, err = modbus.NewClient(&modbus.ClientConfiguration{
		URL: fmt.Sprintf("tcp://%s:%d", "192.168.88.1", 502),
		// URL:     fmt.Sprintf("tcp://%s:%d", "localhost", 10502),
		Timeout: 5 * time.Second,
	})

	if err != nil {
		panic(err.Error())
	}
	client.SetUnitId(1)
	for {
		err = client.Open()
		if err != nil {
			fmt.Printf("client %s\n", err.Error())
			time.Sleep(time.Second)
		} else {
			break
		}
	}
	for {
		time.Sleep(time.Second)
		res, err := client.ReadExceptionStatus()
		if err != nil {
			fmt.Printf("client %s\n", err.Error())
		}
		fmt.Printf("res:=%v \t", res)
		res, err = client.ReportServerID()
		if err != nil {
			fmt.Printf("client %s\n", err.Error())
		}
		fmt.Printf("ids:=%v\n", res)
	}

}
