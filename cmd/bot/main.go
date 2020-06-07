package main

import (
	"flag"
	"image/color"
	"strings"

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
	usersStr := flag.String("users", "", "Telegram usernames of the allowed users (comma-separated)")
	flag.Parse()

	if *debug {
		log.SetLevel(log.DebugLevel)
	}

	if *token == "" {
		log.Fatalf("Please specify Telegram API token via --token")
	}

	users := parseList(*usersStr)
	if len(users) == 0 {
		log.Fatalf("Please specify a list of allowed telegram users via --users")
	}
	log.Debugf("allowed users: %v", users)

	userMap := map[string]bool{}
	for _, user := range users {
		userMap[user] = true
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
		username := u.Message.From.UserName

		if !userMap[username] {
			log.Warnf("unauthorized user: %s", username)
			msg := tg.NewMessage(u.Message.Chat.ID, "I don't know you. Go away!")
			bot.Send(msg)
			continue
		}

		if u.Message == nil { // ignore any non-Message Updates
			continue
		}

		if u.Message.Text != "" {
			log.Debugf("Text message from %s: %s", username, u.Message.Text)
			scr.Clear()

			disp := screen.NewDisplayer(scr)
			tinyfont.WriteLine(disp, &tinyfont.Org01, 0, 5, []byte(u.Message.Text), white)
			err = disp.Display()
			if err != nil {
				log.Fatalf("display error: %v", err)
			}

			msg := tg.NewMessage(u.Message.Chat.ID, "Message delivered")
			bot.Send(msg)
		}

		if u.Message.Photo != nil {
			log.Debugf("Photo message from %s", u.Message.From.UserName)
		}
	}
}

func parseList(str string) []string {
	res := []string{}
	for _, tok := range strings.Split(str, ",") {
		tok := strings.TrimSpace(tok)
		if tok != "" {
			res = append(res, tok)
		}
	}
	return res
}
