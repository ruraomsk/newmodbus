package server

import (
	"fmt"
	"sync"
	"time"

	"github.com/ruraomsk/newmodbus/modbus"
)

var server *modbus.ModbusServer
var err error

type handler struct {
	lock       sync.Mutex
	mapClients map[string]time.Time
}

var eh = &handler{mapClients: make(map[string]time.Time)}

func (h *handler) HandleCoils(req *modbus.CoilsRequest) (res []bool, err error) {
	h.lock.Lock()
	defer h.lock.Unlock()
	err = modbus.ErrIllegalFunction
	return
}

func (h *handler) HandleDiscreteInputs(req *modbus.DiscreteInputsRequest) (res []bool, err error) {
	h.lock.Lock()
	defer h.lock.Unlock()
	err = modbus.ErrIllegalFunction
	return
}

func (h *handler) HandleHoldingRegisters(req *modbus.HoldingRegistersRequest) (res []uint16, err error) {
	h.lock.Lock()
	defer h.lock.Unlock()
	err = modbus.ErrIllegalFunction
	return
}
func (h *handler) HandleInputRegisters(req *modbus.InputRegistersRequest) (res []uint16, err error) {
	h.lock.Lock()
	defer h.lock.Unlock()
	err = modbus.ErrIllegalFunction
	return
}
func (h *handler) HandleReadExceptionStatus(req *modbus.ControlRequest) (res []uint8, err error) {
	h.lock.Lock()
	defer h.lock.Unlock()
	res = []byte{0x01, 0x2}
	fmt.Printf("0x11 %v", res)
	err = nil
	return
}
func (h *handler) HandleReportServerID(req *modbus.ControlRequest) (res []uint8, err error) {
	h.lock.Lock()
	defer h.lock.Unlock()
	res = []byte{0x57}
	fmt.Printf("0x07 %v", res)
	err = nil
	return
}

func GetClients() map[string]time.Time {
	eh.lock.Lock()
	defer eh.lock.Unlock()
	return eh.mapClients
}

func Start() {
	fmt.Println("Server start")
	server, err = modbus.NewServer(&modbus.ServerConfiguration{
		URL:        fmt.Sprintf("tcp://0.0.0.0:10502"),
		Timeout:    30 * time.Second,
		MaxClients: 5,
	}, eh)
	if err != nil {
		fmt.Printf("Не могу создать сервер %v", err)
		return
	}

	err = server.Start()
	if err != nil {
		fmt.Printf("Не могу запустить сервер %v", err)
		return
	}
	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
	}
}
