package main

import (
	"fmt"
	"github.com/FedeDP/golight/backlight"
	"github.com/FedeDP/golight/capture"
	"github.com/FedeDP/golight/day"
	"github.com/FedeDP/golight/dimmer"
	"github.com/FedeDP/golight/dpms"
	"github.com/FedeDP/golight/gamma"
	"github.com/FedeDP/golight/location"
	"github.com/FedeDP/golight/signals"
	"github.com/FedeDP/golight/state"
	"github.com/FedeDP/golight/upower"
)

/*
TODO:
 * Move Clightd wrapper to its own external module
 * Fix: only start gamma once a first location has been received
 * Pass by pointer where needed
 * Fix memleaks (?)
 * Implement conf parsing (https://github.com/spf13/viper)
 * Implement dbus server exposing state
 */

func main() {
	locC := location.Subscribe()
	gammaC := gamma.Subscribe() // gamma
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

	/* Init */
	go gamma.Update()
	go capture.Update(blC)

	quit := false
	for !quit {
		select {
		case v := <-locC:
			location.Update(v)
			gammaC = gamma.Subscribe() // update timer to next event

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
			upower.Update()

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
