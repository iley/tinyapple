package fonts

import (
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/proggy"
)

// Extra characters for proggy bundled with tinyfont.
var proggyExtraGlyphs = []tinyfont.Glyph{
	/* ° */ tinyfont.Glyph{Rune: 176, Width: 0x3, Height: 0x3, XAdvance: 0x6, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0x55, 0x00}},
}

func init() {
	proggy.TinySZ8pt7b.Glyphs = append(proggy.TinySZ8pt7b.Glyphs, proggyExtraGlyphs...)
}
