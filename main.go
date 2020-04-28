package main

import (
	"fmt"
	"log"

	"periph.io/x/periph/conn/physic"
	"periph.io/x/periph/conn/spi"
	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/host"
)

func testSPI() error {
	if _, err := host.Init(); err != nil {
		return err
	}

	p, err := spireg.Open("")
	if err != nil {
		return err
	}
	defer p.Close()

	c, err := p.Connect(physic.MegaHertz, spi.Mode3, 8)
	if err != nil {
		return err
	}

	write := []byte{0x10, 0x00}
	read := make([]byte, len(write))
	if err := c.Tx(write, read); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", read[1:])

	return nil
}

func main() {
	err := testSPI()
	if err != nil {
		log.Fatal(err)
	}
}
