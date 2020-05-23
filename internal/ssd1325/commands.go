package ssd1325

const (
	displayOn        = 0xAF
	displayOff       = 0xAE
	setClock         = 0xB3 // value: upper 4 bits = frequency, lower 4 bits = (divide ratio-1)
	setMultiplex     = 0xA8
	setOffset        = 0xA2 // from 0 to 80 for vertical scrolling
	setStartLine     = 0xA1
	masterConfig     = 0xAD
	setRemap         = 0xA0
	setCurrentFull   = 0x84 + 0x03
	setColAddr       = 0x15 // start, end
	setRowAddr       = 0x75 // start, end
	setContrast      = 0x81 // value (0-0x7f)
	normalDisplay    = 0xA4
	invertDisplay    = 0xA7
	setGrayTable     = 0xB8
	setRowPeriod     = 0xB2
	setPhaseLen      = 0xB1
	setChargeComp    = 0xB4
	chargeCompEnable = 0xB0
	setComLevel      = 0xBE
	setVSL           = 0xBF
	gfxAccel         = 0x23
	drawRect         = 0x24
)
