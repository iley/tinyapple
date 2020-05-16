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

	// TODO: Pick appropriate speed.
	conn, err := port.Connect(10*physic.KiloHertz, spi.Mode0, 8)
	if err != nil {
		return nil, err
	}

	dcPin := gpioreg.ByName(dcPinName)
	if dcPin == nil {
		return nil, fmt.Errorf("could not open GPIO %s", dcPin)
	}

	resetPin := gpioreg.ByName(resetPinName)
	if resetPin == nil {
		return nil, fmt.Errorf("could not open GPIO %s", resetPin)
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

func (s *SSD1325) command(cmd []byte) error {
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
	return nil
}

// Init initializes the SSD1325 chip.
func (s *SSD1325) init() error {
	log.Debugf("initializing SSD1325...")

	err := s.reset()
	if err != nil {
		return err
	}

	err = s.command(
		[]byte{
			0xAE,       // Display off (all pixels off)
			0xA0, 0x53, // Segment remap (com split, com remap, nibble remap, column remap)
			0xA1, 0x00, // Display start line
			0xA2, 0x00, // Display offset
			0xA4,       // regular display
			0xA8, 0x7F, // set multiplex ratio: 127

			0xB8, 0x01, 0x11, // Set greyscale table
			0x22, 0x32, 0x43, // .. cont
			0x54, 0x65, 0x76, // .. cont

			0xB3, 0x00, // Front clock divider: 0, Fosc: 0
			0xAB, 0x01, // Enable Internal Vdd

			0xB1, 0xF1, // Set phase periods - 1: 1 clk, 2: 15 clks
			0xBC, 0x08, // Pre-charge voltage: Vcomh
			0xBE, 0x07, // COM deselect voltage level: 0.86 x Vcc

			0xD5, 0x62, // Enable 2nd pre-charge
			0xB6, 0x0F, // 2nd Pre-charge period: 15 clks
		},
	)
	if err != nil {
		return err
	}

	return nil
}

// Close closes the SPI port.
func (s *SSD1325) Close() error {
	return s.port.Close()
}

// TODO: The rest of methods.
