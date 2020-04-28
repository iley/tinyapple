package ssd1325

import (
	"periph.io/x/periph/conn/physic"
	"periph.io/x/periph/conn/spi"
	"periph.io/x/periph/conn/spi/spireg"
)

// SSD1325 provides support for displays based on SSD1325 chip.
// Only SPI mode is supported.
type SSD1325 struct {
	spiport spi.PortCloser
	conn    spi.Conn
}

// NewSSD1325 creates an initializes a new instance of SSD1325.
func NewSSD1325() (*SSD1325, error) {
	p, err := spireg.Open("")
	if err != nil {
		return nil, err
	}
	c, err := p.Connect(physic.MegaHertz, spi.Mode3, 8)
	if err != nil {
		return nil, err
	}
	s := &SSD1325{
		spiport: p,
		conn:    c,
	}
	err = s.Init()
	if err != nil {
		p.Close()
		return nil, err
	}
	return s, nil
}

// Init initializes the SSD1325 chip.
func (s *SSD1325) Init() error {
	return nil
}

// Close closes the SPI connection to SSD1325 chip.
func (s *SSD1325) Close() error {
	return s.spiport.Close()
}
