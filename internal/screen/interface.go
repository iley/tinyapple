package screen

import (
	"fmt"
	"image"

	log "github.com/sirupsen/logrus"
)

// Interface describes a screen device.
type Interface interface {
	Width() int
	Height() int
	Close() error
	Clear() error
	DrawBuffer(data []byte) error
	PackImage(img image.Image) ([]byte, error)
}

// DrawImage is a utility function for drawing an unpacked image.
func DrawImage(scr Interface, img image.Image) error {
	log.Debugf("drawing an image...")
	buf, err := scr.PackImage(img)
	if err != nil {
		return fmt.Errorf("could not pack image: %w", err)
	}
	return scr.DrawBuffer(buf)
}
