package main

import (
	"flag"
	"time"

	log "github.com/sirupsen/logrus"
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
)

func main() {
	var err error

	pinName := flag.String("pin", "GPIO0", "Name of the pin")
	high := flag.Bool("high", false, "Set pin to high")
	low := flag.Bool("low", false, "Set pin to low")
	delay := flag.Duration("delay", 5*time.Second, "Sleep duration")

	flag.Parse()

	if *high && *low || !*high && !*low {
		log.Fatalf("Please specify ether -high or -low")
	}

	log.Infof("initializing host...")
	if _, err = host.Init(); err != nil {
		log.Fatalf("host initialization error: %v", err)
	}

	pin := gpioreg.ByName(*pinName)
	if pin == nil {
		log.Fatalf("could not open GPIO %s", *pinName)
	}

	log.Infof("setting pin state to %v", *high)
	if *high {
		err = pin.Out(gpio.High)
	} else {
		err = pin.Out(gpio.Low)
	}
	if err != nil {
		log.Fatalf("error writing: %v", err)
	}

	log.Infof("waiting....")
	time.Sleep(*delay)
}
