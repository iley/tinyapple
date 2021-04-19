package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/nfnt/resize"
	log "github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
	"periph.io/x/periph/host"
	"tinygo.org/x/tinyfont"

	"github.com/iley/littlemac/internal/screen"
	"github.com/iley/littlemac/internal/screen/ssd1325"
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

	bot.Handle(tb.OnText, authorize(bot, users, func(msg *tb.Message) {
		handleText(bot, msg, scr)
	}))

	bot.Handle(tb.OnPhoto, authorize(bot, users, func(msg *tb.Message) {
		handlePhoto(bot, msg, scr)
	}))

	bot.Handle(tb.OnDocument, authorize(bot, users, func(msg *tb.Message) {
		handleDocument(bot, msg, scr)
	}))

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

// authorize wraps a handler to add an authorization check based on allowlist.
func authorize(bot *tb.Bot, users []string, handler func(*tb.Message)) func(*tb.Message) {
	userMap := map[string]bool{}
	for _, user := range users {
		userMap[user] = true
	}

	return func(msg *tb.Message) {
		if !userMap[msg.Sender.Username] {
			log.Warnf("unauthorized user: %s", msg.Sender.Username)
			bot.Send(msg.Sender, "I don't know you. Go away!")
			return
		}
		handler(msg)
	}
}

func handleText(bot *tb.Bot, msg *tb.Message, scr screen.Interface) {
	disp := screen.NewDisplayer(scr)
	tinyfont.WriteLine(disp, &tinyfont.Org01, 0, 5, msg.Text, white)
	err := disp.Display()
	if err != nil {
		log.Fatalf("display error: %v", err)
	}
	bot.Send(msg.Sender, "Message delivered")
}

func handlePhoto(bot *tb.Bot, msg *tb.Message, scr screen.Interface) {
	log.Debugf("photo message received")
	if msg.Photo == nil {
		log.Errorf("invalid photo message")
		return
	}
	handleImageFile(bot, msg, &msg.Photo.File, scr)
}

func handleDocument(bot *tb.Bot, msg *tb.Message, scr screen.Interface) {
	if msg.Document == nil {
		log.Errorf("invalid document message")
		return
	}
	handleImageFile(bot, msg, &msg.Document.File, scr)
}

func handleImageFile(bot *tb.Bot, msg *tb.Message, file *tb.File, scr screen.Interface) {
	img, err := fetchImage(bot, file)
	if err != nil {
		log.Errorf("error fetching image: %v", err)
		bot.Send(msg.Sender, fmt.Sprintf("Could not fetch image: %v", err))
		return
	}

	// Scale down the image if it's larger than the screen.
	if imgWidth(img) > scr.Width() || imgHeight(img) > scr.Height() {
		log.Debugf("resizing image from %dx%d to %dx%d", imgWidth(img), imgHeight(img), scr.Width(), scr.Height())
		img = resize.Thumbnail(uint(scr.Width()), uint(scr.Height()), img, resize.Lanczos3)
	}

	// Surround image with black pixels if it's smaller than the screen.
	// This might also be caused by resizing above.
	if imgWidth(img) < scr.Width() || imgHeight(img) < scr.Height() {
		log.Debugf("extending image from %dx%d to %dx%d", imgWidth(img), imgHeight(img), scr.Width(), scr.Height())
		img = extendImage(img, scr.Width(), scr.Height())
	}

	err = screen.DrawImage(scr, img)
	if err != nil {
		log.Errorf("could not draw image: %v", err)
	}
}

func fetchFile(bot *tb.Bot, file *tb.File) ([]byte, error) {
	tmpfile, err := ioutil.TempFile("", "img")
	if err != nil {
		return nil, fmt.Errorf("could not create temporary file: %w", err)
	}

	defer os.Remove(tmpfile.Name())

	if err = tmpfile.Close(); err != nil {
		return nil, fmt.Errorf("could not close temporary file: %w", err)
	}

	if err = bot.Download(file, tmpfile.Name()); err != nil {
		return nil, fmt.Errorf("could not download file: %w", err)
	}

	return ioutil.ReadFile(tmpfile.Name())
}

func fetchImage(bot *tb.Bot, file *tb.File) (image.Image, error) {
	log.Debugf("downloading an image...")
	data, err := fetchFile(bot, file)
	if err != nil {
		return nil, err
	}

	contentType := http.DetectContentType(data)
	log.Debugf("downloaded file of size %d, content type %s", len(data), contentType)

	reader := bytes.NewReader(data)
	if contentType == "image/png" {
		return png.Decode(reader)
	} else if contentType == "image/jpeg" {
		return jpeg.Decode(reader)
	} else {
		return nil, fmt.Errorf("Unsupported image type: %s. Only PNG and JPEG are supported", contentType)
	}
}

// extendImage extends a small image into a fixed size by adding black pixels around
func extendImage(img image.Image, width, height int) image.Image {
	res := image.NewRGBA(image.Rect(0, 0, width, height))
	offsetX := (width - imgWidth(img)) / 2
	offsetY := (height - imgHeight(img)) / 2
	for y := 0; y < imgHeight(img); y++ {
		for x := 0; x < imgWidth(img); x++ {
			res.Set(offsetX+x, offsetY+y, img.At(x, y))
		}
	}
	return res
}

func imgWidth(img image.Image) int {
	return img.Bounds().Max.X
}

func imgHeight(img image.Image) int {
	return img.Bounds().Max.Y
}
