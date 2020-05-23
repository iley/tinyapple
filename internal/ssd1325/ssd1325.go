package ssd1325

// Inspired by https://github.com/adafruit/Adafruit_SSD1325_Library

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/conn/physic"
	"periph.io/x/periph/conn/spi"
	"periph.io/x/periph/conn/spi/spireg"
)

// ScreenWidth is the screen width in pixels
const ScreenWidth = 128

// ScreenHeight is the screen height in pixels
const ScreenHeight = 64

// SSD1325 provides support for displays based on SSD1325 chip.
// Only SPI mode is supported.
type SSD1325 struct {
	port     spi.PortCloser
	conn     spi.Conn
	dcPin    gpio.PinOut
	resetPin gpio.PinOut
}

// New creates an initializes a new instance of SSD1325.
func New(spiDevName, dcPinName, resetPinName string) (*SSD1325, error) {
	port, err := spireg.Open(spiDevName)
	if err != nil {
		return nil, err
	}

	conn, err := port.Connect(4*physic.MegaHertz, spi.Mode0, 8)
	if err != nil {
		return nil, err
	}

	dcPin := gpioreg.ByName(dcPinName)
	if dcPin == nil {
		return nil, fmt.Errorf("could not open GPIO %s", dcPinName)
	}

	resetPin := gpioreg.ByName(resetPinName)
	if resetPin == nil {
		return nil, fmt.Errorf("could not open GPIO %s", resetPinName)
	}

	s := &SSD1325{
		port:     port,
		conn:     conn,
		dcPin:    dcPin,
		resetPin: resetPin,
	}

	err = s.init()
	if err != nil {
		port.Close()
		return nil, err
	}
	return s, nil
}

func (s *SSD1325) spiWrite(dcLevel gpio.Level, data []byte) error {
	err := s.dcPin.Out(dcLevel)
	if err != nil {
		return err
	}
	readBuf := make([]byte, len(data))
	err = s.conn.Tx(data, readBuf)
	return err
}

func (s *SSD1325) data(data ...byte) error {
	return s.spiWrite(gpio.High, data)
}

func (s *SSD1325) command(cmd ...byte) error {
	return s.spiWrite(gpio.Low, cmd)
}

func (s *SSD1325) reset() error {
	log.Debugf("resetting the display module")

	err := s.resetPin.Out(gpio.High)
	if err != nil {
		return err
	}
	time.Sleep(time.Millisecond)

	err = s.resetPin.Out(gpio.Low)
	if err != nil {
		return err
	}
	time.Sleep(10 * time.Millisecond)

	err = s.resetPin.Out(gpio.High)
	if err != nil {
		return err
	}

	// Wait a while after reset.
	time.Sleep(500 * time.Millisecond)

	return nil
}

// Init initializes the SSD1325 chip.
func (s *SSD1325) init() error {
	log.Debugf("initializing SSD1325...")

	err := s.reset()
	if err != nil {
		return err
	}

	err = s.command(0xAF) // Display on.
	if err != nil {
		return err
	}
	time.Sleep(100 * time.Millisecond)

	err = s.command(
		displayOff,
		setClock, 0xF1,
		setMultiplex, 0x3F, // multiplex ratio duty = 1/64
		setOffset, 0x4C,
		setStartLine, 0x00,
		masterConfig, 0x02,
		setRemap, 0x56,
		setCurrentFull,
		setGrayTable, 0x01, 0x11, 0x22, 0x32, 0x43, 0x54, 0x65, 0x76,
		setContrast, 0x7F, // max
		setRowPeriod, 0x51,
		setPhaseLen, 0x55,
		setChargeComp, 0x02,
		chargeCompEnable, 0x28,
		setComLevel, 0x1C,
		setVSL, 0x0D|0x02, // set high voltage level of COM pin.
		normalDisplay,
		gfxAccel, 0x01, // enable drawrect

		displayOn,
	)
	if err != nil {
		return fmt.Errorf("Initialization sequence failed: %w", err)
	}

	log.Debugf("initialization done")
	return nil
}

// Close closes the SPI port.
func (s *SSD1325) Close() error {
	s.reset()
	return s.port.Close()
}

// Display a frame buffer on the screen.
func (s *SSD1325) Display(data []byte) error {
	if len(data) != ScreenWidth*ScreenHeight {
		return fmt.Errorf("wrong size of buffer. expected %v", ScreenWidth*ScreenHeight)
	}

	err := s.command(
		setColAddr, 0x00, 0x3f,
		setRowAddr, 0x00, 0x3f,
	)
	if err != nil {
		return err
	}

	// TODO: Draw whole line at a time.
	for row := 0; row < ScreenHeight; row++ {
		// Pack each two pixels into a byte.
		for col := 0; col < ScreenWidth/2; col += 2 {
			// TODO: Check if this is correct order.
			low := data[row*ScreenWidth+col] & 0x0F
			high := data[row*ScreenWidth+col+1] & 0x0F
			d := (high << 4) | low
			err = s.data(d)
			if err != nil {
				return fmt.Errorf("Error writing pixel %dx%d: %w", row, col, err)
			}
		}
	}
	return nil
}

// DrawRect draws a rectangle on the screen.
func (s *SSD1325) DrawRect(startRow, endRow, startCol, endCol uint, pattern byte) error {
	if startRow >= ScreenHeight || endRow >= ScreenHeight || startRow > endRow {
		return fmt.Errorf("invalid start/end row")
	}
	if startCol >= ScreenWidth || endCol >= ScreenWidth || startCol > endCol {
		return fmt.Errorf("invalid start/end column")
	}
	log.Debugf("drawing a rectangle")
	return s.command(
		drawRect, byte(startRow), byte(endRow), byte(startCol), byte(endCol), pattern,
	)

}
