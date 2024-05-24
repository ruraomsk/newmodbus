package client

import (
	"fmt"
	"time"

	"github.com/ruraomsk/potop/modbus"
)

var client *modbus.ModbusClient
var err error

func Start() {
	fmt.Println("Client start")
	client, err = modbus.NewClient(&modbus.ClientConfiguration{
		URL:     fmt.Sprintf("tcp://%s:%d", "localhost", 10502),
		Timeout: 5 * time.Second,
	})

	if err != nil {
		panic(err.Error())
	}
	client.SetUnitId(1)
	for {
		err = client.Open()
		if err != nil {
			fmt.Println(err.Error())
			time.Sleep(time.Second)
		} else {
			break
		}
	}
	for {
		time.Sleep(time.Second)
		res, err := client.ReadExceptionStatus()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Printf("res:=%v \t", res)
		res, err = client.ReportServerID()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Printf("ids:=%v\n", res)
	}

}
