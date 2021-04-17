package main

import (
	"flag"
	"fmt"
	"image/color"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
	"periph.io/x/periph/host"
	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/proggy"

	"github.com/iley/tinyapple/internal/fonts"
	"github.com/iley/tinyapple/internal/screen"
	"github.com/iley/tinyapple/internal/screen/ssd1325"
)

var (
	timeFont = &fonts.Digits
	dateFont = &proggy.TinySZ8pt7b
)

func main() {
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
	var scr screen.Interface
	scr, err := ssd1325.New(*spiDev, *dcPin, *rstPin)
	if err != nil {
		log.Fatalf("could not initialize screen: %v", err)
	}

	defer scr.Close()

	disp := screen.NewDisplayer(scr)
	white := color.RGBA{0xff, 0xff, 0xff, 255}
	black := color.RGBA{0x00, 0x00, 0x00, 255}

	scrWidth := int16(scr.Width())
	scrHeight := int16(scr.Height())

	for {
		now := time.Now()

		tinydraw.FilledRectangle(disp, 0, 0, scrWidth-1, scrHeight-1, black)

		timeStr := getTimeStr(now)
		tinyfont.WriteLine(disp, timeFont, 10, 32, timeStr, white)

		dateStr := getDateStr(now)
		tinyfont.WriteLine(disp, dateFont, 10, 54, dateStr, white)

		err = disp.Display()

		if err != nil {
			log.Fatalf("display error: %v", err)
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func getTimeStr(now time.Time) string {
	hours, minutes, _ := now.Clock()
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}

func getDateStr(now time.Time) string {
	weekday := time.Now().Weekday().String()
	weekdayShort := weekday[0:3]
	date := now.Format("02 Jan")
	return fmt.Sprintf("%s %s", weekdayShort, date)
}
