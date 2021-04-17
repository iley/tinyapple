package screen

import (
	"image"
	"image/color"

	"tinygo.org/x/drivers"
)

// Displayer implements the Displayer interfaces from TinyGo.
type Displayer struct {
	scr       Interface
	frame     *image.RGBA
	prevFrame *image.RGBA
}

var _ drivers.Displayer = (*Displayer)(nil)

// NewDisplayer creates a new instance of Displayer.
func NewDisplayer(scr Interface) *Displayer {
	return &Displayer{
		scr:       scr,
		frame:     image.NewRGBA(image.Rect(0, 0, scr.Width(), scr.Height())),
		prevFrame: image.NewRGBA(image.Rect(0, 0, scr.Width(), scr.Height())),
	}
}

// Size returns the screen size in pixels.
func (d *Displayer) Size() (int16, int16) {
	return int16(d.scr.Width()), int16(d.scr.Height())
}

// SetPixel sets a pixel value in the in-memory buffer.
func (d *Displayer) SetPixel(x, y int16, c color.RGBA) {
	d.frame.Set(int(x), int(y), c)
}

// Display sends the in-memory buffer to the screen.
func (d *Displayer) Display() error {
	if imagesEqual(d.prevFrame, d.frame) {
		return nil
	}
	err := DrawImage(d.scr, d.frame)
	copy(d.prevFrame.Pix, d.frame.Pix)
	return err
}

func imagesEqual(first, second *image.RGBA) bool {
	// We assume size is always the same.
	for i := range first.Pix {
		if first.Pix[i] != second.Pix[i] {
			return false
		}
	}
	return true
}
