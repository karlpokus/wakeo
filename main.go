package main

import (
	"machine"
	"time"
)

// log to USB serial

func main() {
	serial := machine.Serial
	serial.Configure(machine.UARTConfig{})

	for {
		serial.Write([]byte("hello again from pico\n"))
		time.Sleep(time.Second)
	}
}
