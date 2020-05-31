package screen

import (
	"fmt"
	"image"
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
	buf, err := scr.PackImage(img)
	if err != nil {
		return fmt.Errorf("could not pack image: %w", err)
	}
	return scr.DrawBuffer(buf)
}
