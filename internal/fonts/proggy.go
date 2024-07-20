package fonts

import (
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/proggy"
)

// Extra characters for proggy bundled with tinyfont.
var proggyExtraGlyphs = []tinyfont.Glyph{
	/* ° */ {Rune: 176, Width: 3, Height: 3, XAdvance: 6, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0x55, 0x00}},
	/* ☂ */ {Rune: 9730, Width: 8, Height: 8, XAdvance: 10, XOffset: 2, YOffset: -7, Bitmaps: []uint8{0x08, 0x1c, 0x3e, 0x7f, 0x08, 0x08, 0x0a, 0x04}},
}

func init() {
	proggy.TinySZ8pt7b.Glyphs = append(proggy.TinySZ8pt7b.Glyphs, proggyExtraGlyphs...)
}
