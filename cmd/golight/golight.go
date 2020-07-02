package main

import (
	"fmt"
	"github.com/FedeDP/golight/internal/app/golight/backlight"
	"github.com/FedeDP/golight/internal/app/golight/capture"
	"github.com/FedeDP/golight/internal/app/golight/day"
	"github.com/FedeDP/golight/internal/app/golight/dimmer"
	"github.com/FedeDP/golight/internal/app/golight/dpms"
	"github.com/FedeDP/golight/internal/app/golight/gamma"
	"github.com/FedeDP/golight/internal/app/golight/location"
	"github.com/FedeDP/golight/internal/app/golight/signals"
	"github.com/FedeDP/golight/internal/app/golight/state"
	"github.com/FedeDP/golight/internal/app/golight/upower"
	"time"
)

func main() {
	var gammaC <-chan time.Time
	locC := location.Subscribe()
	sigC := signals.Subscribe() // signal handler
	dayC := day.Subscribe() // just after midnight to compute next day events
	captureC := capture.Subscribe()
	blC := backlight.Subscribe()
	upC := upower.Subscribe()
	dimC := dimmer.Subscribe()
	dpmsC := dpms.Subscribe()

	/* Cleanup functions */
	defer location.Close()
	defer gamma.Close()
	defer capture.Close()
	defer backlight.Close()
	defer upower.Close()
	defer dimmer.Close()
	defer dpms.Close()

	capture.Update(blC)

	quit := false
	for !quit {
		select {
		case v := <-locC:
			location.Update(v)
			firstLoc := state.NextSunrise.IsZero()
			gammaC = gamma.Subscribe() // update timer to next event
			if firstLoc {
				gamma.Update()
			}

		case <-gammaC:
			gammaC = gamma.Subscribe() // update timer to next event
			gamma.Update()

		case <-dayC:
			dayC = day.Subscribe()
			gammaC = gamma.Subscribe() // update timer to next event

		case <-captureC:
			if state.Display == state.DisplayON {
				capture.Update(blC)
			}

		case c := <-blC:
			backlight.Update(c)

		case <-upC:
			if ok, _ := upower.Update(); ok {
                if state.Ac == state.OnBatt {
                    fmt.Println("Current AC state: on Batt.")
                } else {
                    fmt.Println("Current AC state: on AC.")
                }
				/* On new upower state, update all timers */
				dimmer.UpdateTimer()
				dpms.UpdateTimer()
				captureC = capture.Subscribe()
			}

		case d := <-dimC:
			dimmer.Update(d)

		case d := <-dpmsC:
			dpms.Update(d)

		case s := <-sigC:
			fmt.Printf("Received signal %s, quitting.\n", s)
			quit = true
		}
	}
}
