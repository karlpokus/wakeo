package main

import (
	"machine"
	"time"
	"wakeo/pkg/clock"
)

var buildUnix string            // wall-clock time injected at build time in unix ts format
var boot = time.Now()           // monotonic time sampled at boot
var interval = time.Second * 10 // arbitrary interval for checking daytime

func main() {
	// configure LED
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	// parse timestamp as base
	base, err := clock.Parse(buildUnix)
	if err != nil {
		// errors are handled by fast blinking onboard LED
		fatalBlink(led)
		return
	}
	// start timer
	//
	// Blink once if daytime
	ticker := time.NewTicker(interval)
	for range ticker.C {
		t := base.Add(time.Since(boot))
		if clock.IsDay(t) {
			blinkOnce(led)
		}
	}
}

func blinkOnce(led machine.Pin) {
	led.High()
	time.Sleep(time.Second)
	led.Low()
}

func fatalBlink(led machine.Pin) {
	for {
		led.High()
		time.Sleep(100 * time.Millisecond)
		led.Low()
		time.Sleep(100 * time.Millisecond)
	}
}
