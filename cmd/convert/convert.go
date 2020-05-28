package main

import (
	"flag"
	"fmt"

	"github.com/fogleman/gg"
	log "github.com/sirupsen/logrus"

	"github.com/iley/tinyapple/internal/ssd1325"
)

const colors = 16

func main() {
	offsetX := flag.Int("offsetx", 0, "X offset from left")
	offsetY := flag.Int("offsety", 0, "Y offset from top")
	background := flag.Int("background", 0, "background color")
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		log.Fatalf("Usage: convert [FLAGS] PNG")
	}

	pngPath := args[0]

	img, err := gg.LoadPNG(pngPath)
	if err != nil {
		log.Fatalf("could not load input PNG file %s: %v", pngPath, err)
	}

	bounds := img.Bounds()
	if bounds.Max.X > ssd1325.ScreenWidth || bounds.Max.Y > ssd1325.ScreenHeight {
		log.Fatalf("invalid image size %dx%d, must be <= %dx%d", bounds.Max.X, bounds.Max.Y, ssd1325.ScreenWidth, ssd1325.ScreenHeight)
	}

	buf := [ssd1325.ScreenWidth][ssd1325.ScreenHeight]uint8{}
	for y := 0; y < ssd1325.ScreenHeight; y++ {
		for x := 0; x < ssd1325.ScreenWidth; x++ {
			buf[x][y] = uint8(*background)
		}
	}

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			col := uint8((colors - 1) * float64(r+g+b) / (0xffff * 3))
			buf[*offsetX+x][*offsetY+y] = col
		}
	}

	for y := 0; y < ssd1325.ScreenHeight; y++ {
		fmt.Printf("\"")
		for x := 0; x < ssd1325.ScreenWidth; x++ {
			fmt.Printf("\\x%02x", buf[x][y])
		}
		if y == ssd1325.ScreenHeight-1 {
			fmt.Printf("\",\n")
		} else {
			fmt.Printf("\" + \n")
		}
	}
}
