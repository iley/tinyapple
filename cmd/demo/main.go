package main

import (
	"encoding/base64"
	"flag"
	"image/png"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/fogleman/gg"
	log "github.com/sirupsen/logrus"
	"periph.io/x/periph/host"

	"github.com/iley/tinyapple/internal/ssd1325"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	spiDev := flag.String("spi", "/dev/spidev0.1", "path to the SPI device")
	dcPin := flag.String("dc", "GPIO1", "name of the D/C GPIO pin")
	rstPin := flag.String("reset", "GPIO0", "name of the reset GPIO pin")
	debug := flag.Bool("debug", false, "enable debug logging")
	flag.Parse()

	if *debug {
		log.SetLevel(log.DebugLevel)
	}

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		done <- true
	}()

	log.Debugf("initializing host...")
	if _, err := host.Init(); err != nil {
		log.Fatalf("host initialization error: %v", err)
	}

	log.Debugf("initializing display...")
	scr, err := ssd1325.New(*spiDev, *dcPin, *rstPin)
	if err != nil {
		log.Fatalf("could not initialize screen: %v", err)
	}

	defer scr.Close()

	log.Debugf("starting demo...")

	log.Debugf("circle")
	dc := gg.NewContext(ssd1325.ScreenWidth, ssd1325.ScreenHeight)
	dc.DrawCircle(ssd1325.ScreenWidth/2, ssd1325.ScreenHeight/2, ssd1325.ScreenHeight/2)
	dc.SetRGB(1, 1, 1)
	dc.Stroke()
	img := dc.Image()
	err = scr.DrawImage(img)
	if err != nil {
		log.Fatalf("could not draw image: %v", err)
	}
	time.Sleep(5 * time.Second)

	log.Debugf("animation")
	packedFrames := make([][]byte, len(frames))
	for i := range frames {
		pngData := base64.NewDecoder(base64.StdEncoding, strings.NewReader(frames[i]))
		img, err := png.Decode(pngData)
		if err != nil {
			log.Errorf("could not decode PNG: %v", err)
		}
		packedFrames[i], err = ssd1325.PackImage(img)
		if err != nil {
			log.Fatalf("cannot pack frame: %v", err)
		}
	}

mainloop:
	for {
		select {
		case <-done:
			log.Infof("exiting...")
			break mainloop
		default:
		}
		for _, frame := range packedFrames {
			err = scr.DrawBufferPacked(frame)
			if err != nil {
				log.Fatalf("cannot draw frame: %v", err)
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}
