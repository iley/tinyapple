package main

import (
	"flag"
	"image/color"

	tg "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
	"periph.io/x/periph/host"
	"tinygo.org/x/tinyfont"

	"github.com/iley/tinyapple/internal/screen"
	"github.com/iley/tinyapple/internal/screen/ssd1325"
)

var white = color.RGBA{0xff, 0xff, 0xff, 255}

func main() {
	spiDev := flag.String("spi", "/dev/spidev0.1", "path to the SPI device")
	dcPin := flag.String("dc", "GPIO1", "name of the D/C GPIO pin")
	rstPin := flag.String("reset", "GPIO0", "name of the reset GPIO pin")
	debug := flag.Bool("debug", false, "enable debug logging")
	token := flag.String("token", "", "Telegram API token")
	flag.Parse()

	if *debug {
		log.SetLevel(log.DebugLevel)
	}

	if *token == "" {
		log.Fatalf("Please specify Telegram API token via --token")
	}

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

	bot, err := tg.NewBotAPI(*token)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = *debug

	log.Debugf("Authorized on account %s", bot.Self.UserName)

	upd := tg.NewUpdate(0)
	upd.Timeout = 60

	updChan, err := bot.GetUpdatesChan(upd)

	for u := range updChan {
		if u.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Debugf("[%s] %s", u.Message.From.UserName, u.Message.Text)
		scr.Clear()

		disp := screen.NewDisplayer(scr)
		tinyfont.WriteLine(disp, &tinyfont.Org01, 0, 5, []byte(u.Message.Text), white)
		err = disp.Display()
		if err != nil {
			log.Fatalf("display error: %v", err)
		}

		msg := tg.NewMessage(u.Message.Chat.ID, "message delivered")
		bot.Send(msg)
	}
}
