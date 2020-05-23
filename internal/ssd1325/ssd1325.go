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

	// conn, err := port.Connect(4*physic.MegaHertz, spi.Mode0|spi.HalfDuplex, 8)
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

func (s *SSD1325) data(data []byte) error {
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
		0xAE,       // Display off.
		0xB3, 0xF1, // Set clock divider.
		0xA8, 0x3F, // Set multiplex ratio duty = 1/64
		0xA2, 0x4C, // Set display offset.
		0xA1, 0x00, // Set display start line.
		0xAD, 0x02, // Master config.
		0xA0, 0x56, // Segment remap.

		0x84+0x2, // Set full current range.

		0xB8, 0x01, 0x11, // Set greyscale table
		0x22, 0x32, 0x43, // ...
		0x54, 0x65, 0x76, // ...

		0x81, 0x7F, // Set contrast to max.
		0xB2, 0x51, // Set row period.
		0xB1, 0x55, // Set phase len.
		0xB4, 0x02, // Set charge comp.
		0xB0, 0x28, // Charge comp. enable.
		0xBE, 0x1C,
		0xBF, 0x0D|0x02, // Set high voltage level of COM pin.

		0xA4, // Normal display mode.
		0xAF, // Display on.
	)
	if err != nil {
		return err
	}

	log.Debugf("initialization done")
	return nil
}

// Close closes the SPI port.
func (s *SSD1325) Close() error {
	return s.port.Close()
}

// Display a frame buffer on the screen.
func (s *SSD1325) Display(data []byte) error {
	if len(data) != ScreenWidth*ScreenHeight {
		return fmt.Errorf("wrong size of buffer. expected %v", ScreenWidth*ScreenHeight)
	}

	err := s.command(
		0x15, // Set column address.
		0x00, // Set column start address.
		0x3f, // Set column end address.
		0x75, // Set row address.
		0x00, // Set row start address.
		0x3f, // Set row end address.
	)
	if err != nil {
		return err
	}

	for row := 0; row < ScreenHeight; row++ {
		rowStart := row * ScreenWidth
		err = s.data(data[rowStart : rowStart+ScreenWidth])
		if err != nil {
			return err
		}
	}
	return nil
}
