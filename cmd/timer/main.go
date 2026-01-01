package main

import (
	"machine"
	"time"
)

// Blink onboard LED on an interval
// like 10s during development

var interval = time.Second * 10

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	ticker := time.NewTicker(interval)

	for {
		<-ticker.C
		blinkOnce(led)
	}
}

func blinkOnce(led machine.Pin) {
	led.High()
	time.Sleep(time.Second)
	led.Low()
}
