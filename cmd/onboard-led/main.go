package main

import (
	"machine"
	"time"
)

// blink onboard LED

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		// led.Toggle()
		led.High()
		time.Sleep(time.Second)

		led.Low()
		time.Sleep(time.Second)
	}
}
