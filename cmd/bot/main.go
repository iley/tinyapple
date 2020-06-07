package main

import (
	"flag"
	"image/color"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
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
	scr, err := ssd1325.New(*spiDev, *dcPin, *rstPin)
	if err != nil {
		log.Fatalf("could not initialize screen: %v", err)
	}

	bot, err := tb.NewBot(tb.Settings{
		Token:  *token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	bot.Handle(tb.OnText, func(msg *tb.Message) {
		if !userMap[msg.Sender.Username] {
			log.Warnf("unauthorized user: %s", msg.Sender.Username)
			bot.Send(msg.Sender, "I don't know you. Go away!")
			return
		}
		handleTextMessage(bot, msg, scr)
	})

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		log.Infof("exiting...")
		// Gracefully shutdown.
		bot.Stop()
	}()

	bot.Start()
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

func handleTextMessage(bot *tb.Bot, msg *tb.Message, scr screen.Interface) {
	disp := screen.NewDisplayer(scr)
	tinyfont.WriteLine(disp, &tinyfont.Org01, 0, 5, []byte(msg.Text), white)
	err := disp.Display()
	if err != nil {
		log.Fatalf("display error: %v", err)
	}

	bot.Send(msg.Sender, "Message delivered")
}
