package freesans

import (
	"tinygo.org/x/tinyfont"
)

var Regular12pt7b = tinyfont.Font{
	Glyphs: []tinyfont.Glyph{
		/*   */ tinyfont.Glyph{Rune: 32, Width: 0x0, Height: 0x0, XAdvance: 0x6, XOffset: 0, YOffset: 1, Bitmaps: []uint8{}},
		/* ! */ tinyfont.Glyph{Rune: 33, Width: 0x2, Height: 0x12, XAdvance: 0x8, XOffset: 3, YOffset: -17, Bitmaps: []uint8{0xff, 0xff, 0xff, 0xf0, 0xf0}},
		/* " */ tinyfont.Glyph{Rune: 34, Width: 0x6, Height: 0x6, XAdvance: 0x8, XOffset: 1, YOffset: -16, Bitmaps: []uint8{0xcf, 0x3c, 0xf3, 0x8a, 0x20}},
		/* # */ tinyfont.Glyph{Rune: 35, Width: 0xd, Height: 0x10, XAdvance: 0xd, XOffset: 0, YOffset: -15, Bitmaps: []uint8{0x6, 0x30, 0x31, 0x3, 0x18, 0x18, 0xc7, 0xff, 0xbf, 0xfc, 0x31, 0x3, 0x18, 0x18, 0xc7, 0xff, 0xbf, 0xfc, 0x31, 0x1, 0x18, 0x18, 0xc0, 0xc6, 0x6, 0x30}},
		/* $ */ tinyfont.Glyph{Rune: 36, Width: 0xb, Height: 0x14, XAdvance: 0xd, XOffset: 1, YOffset: -17, Bitmaps: []uint8{0x4, 0x3, 0xe1, 0xff, 0x72, 0x6c, 0x47, 0x88, 0xf1, 0x7, 0x20, 0x7e, 0x3, 0xf0, 0x17, 0x2, 0x3c, 0x47, 0x88, 0xf1, 0x1b, 0x26, 0x7f, 0xc3, 0xe0, 0x10, 0x2, 0x0}},
		/* % */ tinyfont.Glyph{Rune: 37, Width: 0x14, Height: 0x11, XAdvance: 0x15, XOffset: 1, YOffset: -16, Bitmaps: []uint8{0x0, 0x6, 0x3, 0xc0, 0x40, 0x7e, 0xc, 0xe, 0x70, 0x80, 0xc3, 0x18, 0xc, 0x31, 0x0, 0xe7, 0x30, 0x7, 0xe6, 0x0, 0x3c, 0x40, 0x0, 0xc, 0x7c, 0x0, 0x8f, 0xe0, 0x19, 0xc7, 0x1, 0x18, 0x30, 0x31, 0x83, 0x2, 0x1c, 0x70, 0x40, 0xfe, 0x4, 0x7, 0xc0}},
		/* & */ tinyfont.Glyph{Rune: 38, Width: 0xe, Height: 0x11, XAdvance: 0x10, XOffset: 1, YOffset: -16, Bitmaps: []uint8{0xf, 0x0, 0x7e, 0x3, 0x9c, 0xc, 0x30, 0x30, 0xc0, 0xe7, 0x1, 0xf8, 0x3, 0x80, 0x3e, 0x1, 0xcc, 0x6e, 0x19, 0xb0, 0x7c, 0xc0, 0xf3, 0x3, 0xce, 0x1f, 0x9f, 0xe6, 0x1e, 0x1c}},
		/* ' */ tinyfont.Glyph{Rune: 39, Width: 0x2, Height: 0x6, XAdvance: 0x5, XOffset: 1, YOffset: -16, Bitmaps: []uint8{0xff, 0xa0}},
		/* ( */ tinyfont.Glyph{Rune: 40, Width: 0x5, Height: 0x17, XAdvance: 0x8, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0x8, 0x8c, 0x66, 0x31, 0x98, 0xc6, 0x31, 0x8c, 0x63, 0x8, 0x63, 0x8, 0x61, 0xc, 0x20}},
		/* ) */ tinyfont.Glyph{Rune: 41, Width: 0x5, Height: 0x17, XAdvance: 0x8, XOffset: 1, YOffset: -17, Bitmaps: []uint8{0x82, 0x18, 0xc3, 0x18, 0xc3, 0x18, 0xc6, 0x31, 0x8c, 0x62, 0x31, 0x88, 0xc4, 0x62, 0x0}},
		/* * */ tinyfont.Glyph{Rune: 42, Width: 0x7, Height: 0x7, XAdvance: 0x9, XOffset: 1, YOffset: -17, Bitmaps: []uint8{0x10, 0x23, 0x5b, 0xe3, 0x8d, 0x91, 0x0}},
		/* + */ tinyfont.Glyph{Rune: 43, Width: 0xa, Height: 0xb, XAdvance: 0xe, XOffset: 2, YOffset: -10, Bitmaps: []uint8{0xc, 0x3, 0x0, 0xc0, 0x30, 0xff, 0xff, 0xf0, 0xc0, 0x30, 0xc, 0x3, 0x0, 0xc0}},
		/* , */ tinyfont.Glyph{Rune: 44, Width: 0x2, Height: 0x6, XAdvance: 0x7, XOffset: 2, YOffset: -1, Bitmaps: []uint8{0xf5, 0x60}},
		/* - */ tinyfont.Glyph{Rune: 45, Width: 0x6, Height: 0x2, XAdvance: 0x8, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0xff, 0xf0}},
		/* . */ tinyfont.Glyph{Rune: 46, Width: 0x2, Height: 0x2, XAdvance: 0x6, XOffset: 2, YOffset: -1, Bitmaps: []uint8{0xf0}},
		/* / */ tinyfont.Glyph{Rune: 47, Width: 0x7, Height: 0x12, XAdvance: 0x7, XOffset: 0, YOffset: -17, Bitmaps: []uint8{0x2, 0xc, 0x10, 0x20, 0xc1, 0x2, 0xc, 0x10, 0x20, 0xc1, 0x2, 0xc, 0x10, 0x20, 0xc1, 0x0}},
		/* 0 */ tinyfont.Glyph{Rune: 48, Width: 0xb, Height: 0x11, XAdvance: 0xd, XOffset: 1, YOffset: -16, Bitmaps: []uint8{0x1f, 0x7, 0xf1, 0xc7, 0x30, 0x6e, 0xf, 0x80, 0xf0, 0x1e, 0x3, 0xc0, 0x78, 0xf, 0x1, 0xe0, 0x3c, 0xe, 0xc1, 0x9c, 0x71, 0xfc, 0x1f, 0x0}},
		/* 1 */ tinyfont.Glyph{Rune: 49, Width: 0x5, Height: 0x11, XAdvance: 0xd, XOffset: 3, YOffset: -16, Bitmaps: []uint8{0x8, 0xcf, 0xff, 0x8c, 0x63, 0x18, 0xc6, 0x31, 0x8c, 0x63, 0x18}},
		/* 2 */ tinyfont.Glyph{Rune: 50, Width: 0xb, Height: 0x11, XAdvance: 0xd, XOffset: 1, YOffset: -16, Bitmaps: []uint8{0x1f, 0xf, 0xf9, 0x87, 0x60, 0x7c, 0x6, 0x0, 0xc0, 0x18, 0x7, 0x1, 0xc0, 0xf0, 0x78, 0x1c, 0x6, 0x0, 0x80, 0x30, 0x7, 0xff, 0xff, 0xe0}},
		/* 3 */ tinyfont.Glyph{Rune: 51, Width: 0xb, Height: 0x11, XAdvance: 0xd, XOffset: 1, YOffset: -16, Bitmaps: []uint8{0x3f, 0xf, 0xf3, 0x87, 0x60, 0x6c, 0xc, 0x1, 0x80, 0x70, 0x7c, 0xf, 0x80, 0x18, 0x1, 0x80, 0x3c, 0x7, 0x80, 0xd8, 0x73, 0xfc, 0x1f, 0x0}},
		/* 4 */ tinyfont.Glyph{Rune: 52, Width: 0xb, Height: 0x11, XAdvance: 0xd, XOffset: 1, YOffset: -16, Bitmaps: []uint8{0x1, 0x80, 0x70, 0xe, 0x3, 0xc0, 0xd8, 0x1b, 0x6, 0x61, 0x8c, 0x21, 0x8c, 0x33, 0x6, 0x7f, 0xff, 0xfe, 0x3, 0x0, 0x60, 0xc, 0x1, 0x80}},
		/* 5 */ tinyfont.Glyph{Rune: 53, Width: 0xb, Height: 0x11, XAdvance: 0xd, XOffset: 1, YOffset: -16, Bitmaps: []uint8{0x3f, 0xcf, 0xf9, 0x80, 0x30, 0x6, 0x0, 0xde, 0x1f, 0xe7, 0xe, 0x0, 0xe0, 0xc, 0x1, 0x80, 0x30, 0x7, 0x81, 0xf8, 0x73, 0xfc, 0x1f, 0x0}},
		/* 6 */ tinyfont.Glyph{Rune: 54, Width: 0xb, Height: 0x11, XAdvance: 0xd, XOffset: 1, YOffset: -16, Bitmaps: []uint8{0xf, 0x7, 0xf9, 0xc3, 0x30, 0x74, 0x1, 0x80, 0x33, 0xc7, 0xfe, 0xf0, 0xdc, 0x1f, 0x1, 0xe0, 0x3c, 0x6, 0xc1, 0xdc, 0x71, 0xfc, 0x1f, 0x0}},
		/* 7 */ tinyfont.Glyph{Rune: 55, Width: 0xb, Height: 0x11, XAdvance: 0xd, XOffset: 1, YOffset: -16, Bitmaps: []uint8{0xff, 0xff, 0xfc, 0x1, 0x0, 0x60, 0x18, 0x2, 0x0, 0xc0, 0x30, 0x6, 0x1, 0x80, 0x30, 0x4, 0x1, 0x80, 0x30, 0x6, 0x1, 0x80, 0x30, 0x0}},
		/* 8 */ tinyfont.Glyph{Rune: 56, Width: 0xb, Height: 0x11, XAdvance: 0xd, XOffset: 1, YOffset: -16, Bitmaps: []uint8{0x1f, 0x7, 0xf1, 0xc7, 0x30, 0x66, 0xc, 0xc1, 0x8c, 0x61, 0xfc, 0x3f, 0x8e, 0x3b, 0x1, 0xe0, 0x3c, 0x7, 0x80, 0xd8, 0x31, 0xfc, 0x1f, 0x0}},
		/* 9 */ tinyfont.Glyph{Rune: 57, Width: 0xb, Height: 0x11, XAdvance: 0xd, XOffset: 1, YOffset: -16, Bitmaps: []uint8{0x1f, 0x7, 0xf1, 0xc7, 0x70, 0x6c, 0x7, 0x80, 0xf0, 0x1e, 0x7, 0x61, 0xef, 0xfc, 0x79, 0x80, 0x30, 0x5, 0x81, 0x98, 0x73, 0xfc, 0x1e, 0x0}},
		/* : */ tinyfont.Glyph{Rune: 58, Width: 0x2, Height: 0xd, XAdvance: 0x6, XOffset: 2, YOffset: -12, Bitmaps: []uint8{0xf0, 0x0, 0x3, 0xc0}},
		/* ; */ tinyfont.Glyph{Rune: 59, Width: 0x2, Height: 0x10, XAdvance: 0x6, XOffset: 2, YOffset: -11, Bitmaps: []uint8{0xf0, 0x0, 0xf, 0x56}},
		/* < */ tinyfont.Glyph{Rune: 60, Width: 0xc, Height: 0xc, XAdvance: 0xe, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0x0, 0x0, 0x7, 0x1, 0xe0, 0xf8, 0x3c, 0xf, 0x0, 0xe0, 0x7, 0xc0, 0xf, 0x0, 0x3c, 0x0, 0xf0, 0x1}},
		/* = */ tinyfont.Glyph{Rune: 61, Width: 0xc, Height: 0x6, XAdvance: 0xe, XOffset: 1, YOffset: -8, Bitmaps: []uint8{0xff, 0xff, 0xff, 0x0, 0x0, 0x0, 0xff, 0xff, 0xff}},
		/* > */ tinyfont.Glyph{Rune: 62, Width: 0xc, Height: 0xc, XAdvance: 0xe, XOffset: 1, YOffset: -11, Bitmaps: []uint8{0x0, 0xe, 0x0, 0x78, 0x1, 0xf0, 0x7, 0xc0, 0xf, 0x0, 0x70, 0x1e, 0xf, 0x3, 0xc0, 0xf0, 0x8, 0x0}},
		/* ? */ tinyfont.Glyph{Rune: 63, Width: 0xa, Height: 0x12, XAdvance: 0xd, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0x1f, 0x1f, 0xee, 0x1b, 0x3, 0xc0, 0xc0, 0x30, 0xc, 0x6, 0x3, 0x81, 0xc0, 0xe0, 0x30, 0xc, 0x3, 0x0, 0x0, 0x0, 0xc, 0x3, 0x0}},
		/* @ */ tinyfont.Glyph{Rune: 64, Width: 0x16, Height: 0x15, XAdvance: 0x18, XOffset: 1, YOffset: -17, Bitmaps: []uint8{0x0, 0xfe, 0x0, 0xf, 0xfe, 0x0, 0xf0, 0x3e, 0x7, 0x0, 0x3c, 0x38, 0x0, 0x30, 0xc1, 0xe0, 0x66, 0xf, 0xd9, 0xd8, 0x61, 0xc3, 0xc3, 0x7, 0xf, 0x1c, 0x1c, 0x3c, 0x60, 0x60, 0xf1, 0x81, 0x83, 0xc6, 0x6, 0x1b, 0x18, 0x38, 0xee, 0x71, 0xe7, 0x18, 0xfd, 0xf8, 0x71, 0xe7, 0xc0, 0xe0, 0x0, 0x1, 0xe0, 0x0, 0x1, 0xff, 0xc0, 0x1, 0xfc, 0x0}},
		/* A */ tinyfont.Glyph{Rune: 65, Width: 0x10, Height: 0x12, XAdvance: 0x10, XOffset: 0, YOffset: -17, Bitmaps: []uint8{0x3, 0xc0, 0x3, 0xc0, 0x3, 0xc0, 0x7, 0xe0, 0x6, 0x60, 0x6, 0x60, 0xe, 0x70, 0xc, 0x30, 0xc, 0x30, 0x1c, 0x38, 0x18, 0x18, 0x1f, 0xf8, 0x3f, 0xfc, 0x30, 0x1c, 0x30, 0xc, 0x70, 0xe, 0x60, 0x6, 0x60, 0x6}},
		/* B */ tinyfont.Glyph{Rune: 66, Width: 0xd, Height: 0x12, XAdvance: 0x10, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0xff, 0xc7, 0xff, 0x30, 0x19, 0x80, 0x6c, 0x3, 0x60, 0x1b, 0x0, 0xd8, 0xc, 0xff, 0xc7, 0xff, 0x30, 0xd, 0x80, 0x3c, 0x1, 0xe0, 0xf, 0x0, 0x78, 0x6, 0xff, 0xf7, 0xfe, 0x0}},
		/* C */ tinyfont.Glyph{Rune: 67, Width: 0xf, Height: 0x12, XAdvance: 0x11, XOffset: 1, YOffset: -17, Bitmaps: []uint8{0x7, 0xe0, 0x3f, 0xf0, 0xe0, 0x73, 0x80, 0x66, 0x0, 0x6c, 0x0, 0x30, 0x0, 0x60, 0x0, 0xc0, 0x1, 0x80, 0x3, 0x0, 0x6, 0x0, 0x6, 0x0, 0x6c, 0x0, 0xdc, 0x3, 0x1e, 0xe, 0x1f, 0xf8, 0xf, 0xc0}},
		/* D */ tinyfont.Glyph{Rune: 68, Width: 0xe, Height: 0x12, XAdvance: 0x11, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0xff, 0x83, 0xff, 0x8c, 0x7, 0x30, 0xe, 0xc0, 0x1b, 0x0, 0x7c, 0x0, 0xf0, 0x3, 0xc0, 0xf, 0x0, 0x3c, 0x0, 0xf0, 0x3, 0xc0, 0x1f, 0x0, 0x6c, 0x3, 0xb0, 0x1c, 0xff, 0xe3, 0xff, 0x0}},
		/* E */ tinyfont.Glyph{Rune: 69, Width: 0xc, Height: 0x12, XAdvance: 0xf, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0xff, 0xff, 0xff, 0xc0, 0xc, 0x0, 0xc0, 0xc, 0x0, 0xc0, 0xc, 0x0, 0xff, 0xef, 0xfe, 0xc0, 0xc, 0x0, 0xc0, 0xc, 0x0, 0xc0, 0xc, 0x0, 0xff, 0xff, 0xff}},
		/* F */ tinyfont.Glyph{Rune: 70, Width: 0xb, Height: 0x12, XAdvance: 0xe, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0xff, 0xff, 0xff, 0x0, 0x60, 0xc, 0x1, 0x80, 0x30, 0x6, 0x0, 0xff, 0xdf, 0xfb, 0x0, 0x60, 0xc, 0x1, 0x80, 0x30, 0x6, 0x0, 0xc0, 0x18, 0x0}},
		/* G */ tinyfont.Glyph{Rune: 71, Width: 0x10, Height: 0x12, XAdvance: 0x12, XOffset: 1, YOffset: -17, Bitmaps: []uint8{0x7, 0xf0, 0x1f, 0xfc, 0x3c, 0x1e, 0x70, 0x6, 0x60, 0x3, 0xe0, 0x0, 0xc0, 0x0, 0xc0, 0x0, 0xc0, 0x7f, 0xc0, 0x7f, 0xc0, 0x3, 0xc0, 0x3, 0x60, 0x3, 0x60, 0x7, 0x30, 0xf, 0x3c, 0x1f, 0x1f, 0xfb, 0x7, 0xe1}},
		/* H */ tinyfont.Glyph{Rune: 72, Width: 0xd, Height: 0x12, XAdvance: 0x11, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0xc0, 0x1e, 0x0, 0xf0, 0x7, 0x80, 0x3c, 0x1, 0xe0, 0xf, 0x0, 0x78, 0x3, 0xff, 0xff, 0xff, 0xf0, 0x7, 0x80, 0x3c, 0x1, 0xe0, 0xf, 0x0, 0x78, 0x3, 0xc0, 0x1e, 0x0, 0xc0}},
		/* I */ tinyfont.Glyph{Rune: 73, Width: 0x2, Height: 0x12, XAdvance: 0x7, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0xff, 0xff, 0xff, 0xff, 0xf0}},
		/* J */ tinyfont.Glyph{Rune: 74, Width: 0x9, Height: 0x12, XAdvance: 0xd, XOffset: 1, YOffset: -17, Bitmaps: []uint8{0x1, 0x80, 0xc0, 0x60, 0x30, 0x18, 0xc, 0x6, 0x3, 0x1, 0x80, 0xc0, 0x60, 0x3c, 0x1e, 0xf, 0x7, 0xc7, 0x7f, 0x1f, 0x0}},
		/* K */ tinyfont.Glyph{Rune: 75, Width: 0xe, Height: 0x12, XAdvance: 0x10, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0xc0, 0x3b, 0x1, 0xcc, 0xe, 0x30, 0x70, 0xc3, 0x83, 0x1c, 0xc, 0xe0, 0x33, 0x80, 0xde, 0x3, 0xdc, 0xe, 0x38, 0x30, 0x60, 0xc1, 0xc3, 0x3, 0x8c, 0x6, 0x30, 0x1c, 0xc0, 0x3b, 0x0, 0x60}},
		/* L */ tinyfont.Glyph{Rune: 76, Width: 0xa, Height: 0x12, XAdvance: 0xe, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0xc0, 0x30, 0xc, 0x3, 0x0, 0xc0, 0x30, 0xc, 0x3, 0x0, 0xc0, 0x30, 0xc, 0x3, 0x0, 0xc0, 0x30, 0xc, 0x3, 0x0, 0xff, 0xff, 0xf0}},
		/* M */ tinyfont.Glyph{Rune: 77, Width: 0x10, Height: 0x12, XAdvance: 0x14, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0xe0, 0x7, 0xe0, 0x7, 0xf0, 0xf, 0xf0, 0xf, 0xd0, 0xf, 0xd8, 0x1b, 0xd8, 0x1b, 0xd8, 0x1b, 0xcc, 0x33, 0xcc, 0x33, 0xcc, 0x33, 0xc6, 0x63, 0xc6, 0x63, 0xc6, 0x63, 0xc3, 0xc3, 0xc3, 0xc3, 0xc3, 0xc3, 0xc1, 0x83}},
		/* N */ tinyfont.Glyph{Rune: 78, Width: 0xd, Height: 0x12, XAdvance: 0x12, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0xe0, 0x1f, 0x0, 0xfc, 0x7, 0xe0, 0x3d, 0x81, 0xee, 0xf, 0x30, 0x79, 0xc3, 0xc6, 0x1e, 0x18, 0xf0, 0xe7, 0x83, 0x3c, 0x1d, 0xe0, 0x6f, 0x1, 0xf8, 0xf, 0xc0, 0x3e, 0x1, 0xc0}},
		/* O */ tinyfont.Glyph{Rune: 79, Width: 0x11, Height: 0x12, XAdvance: 0x13, XOffset: 1, YOffset: -17, Bitmaps: []uint8{0x3, 0xe0, 0xf, 0xfc, 0xf, 0x7, 0x86, 0x0, 0xc6, 0x0, 0x33, 0x0, 0x1b, 0x0, 0x7, 0x80, 0x3, 0xc0, 0x1, 0xe0, 0x0, 0xf0, 0x0, 0x78, 0x0, 0x36, 0x0, 0x33, 0x0, 0x18, 0xc0, 0x18, 0x78, 0x3c, 0x1f, 0xfc, 0x3, 0xf8, 0x0}},
		/* P */ tinyfont.Glyph{Rune: 80, Width: 0xc, Height: 0x12, XAdvance: 0x10, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0xff, 0x8f, 0xfe, 0xc0, 0x6c, 0x3, 0xc0, 0x3c, 0x3, 0xc0, 0x3c, 0x7, 0xff, 0xef, 0xfc, 0xc0, 0xc, 0x0, 0xc0, 0xc, 0x0, 0xc0, 0xc, 0x0, 0xc0, 0xc, 0x0}},
		/* Q */ tinyfont.Glyph{Rune: 81, Width: 0x11, Height: 0x13, XAdvance: 0x13, XOffset: 1, YOffset: -17, Bitmaps: []uint8{0x3, 0xe0, 0xf, 0xfc, 0xf, 0x7, 0x86, 0x0, 0xc6, 0x0, 0x33, 0x0, 0x1b, 0x0, 0x7, 0x80, 0x3, 0xc0, 0x1, 0xe0, 0x0, 0xf0, 0x0, 0x78, 0x0, 0x36, 0x0, 0x33, 0x1, 0x98, 0xc0, 0xfc, 0x78, 0x3c, 0x1f, 0xff, 0x3, 0xf9, 0x80, 0x0, 0x40}},
		/* R */ tinyfont.Glyph{Rune: 82, Width: 0xe, Height: 0x12, XAdvance: 0x11, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0xff, 0xc3, 0xff, 0xcc, 0x3, 0xb0, 0x6, 0xc0, 0x1b, 0x0, 0x6c, 0x1, 0xb0, 0xc, 0xff, 0xe3, 0xff, 0xcc, 0x3, 0xb0, 0x6, 0xc0, 0x1b, 0x0, 0x6c, 0x1, 0xb0, 0x6, 0xc0, 0x1b, 0x0, 0x70}},
		/* S */ tinyfont.Glyph{Rune: 83, Width: 0xe, Height: 0x12, XAdvance: 0x10, XOffset: 1, YOffset: -17, Bitmaps: []uint8{0xf, 0xe0, 0x7f, 0xc3, 0x83, 0x9c, 0x7, 0x60, 0xd, 0x80, 0x6, 0x0, 0x1e, 0x0, 0x3f, 0x80, 0x3f, 0xc0, 0xf, 0x80, 0x7, 0xc0, 0xf, 0x0, 0x3e, 0x0, 0xde, 0xe, 0x3f, 0xf0, 0x3f, 0x80}},
		/* T */ tinyfont.Glyph{Rune: 84, Width: 0xc, Height: 0x12, XAdvance: 0xf, XOffset: 1, YOffset: -17, Bitmaps: []uint8{0xff, 0xff, 0xff, 0x6, 0x0, 0x60, 0x6, 0x0, 0x60, 0x6, 0x0, 0x60, 0x6, 0x0, 0x60, 0x6, 0x0, 0x60, 0x6, 0x0, 0x60, 0x6, 0x0, 0x60, 0x6, 0x0, 0x60}},
		/* U */ tinyfont.Glyph{Rune: 85, Width: 0xd, Height: 0x12, XAdvance: 0x11, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0xc0, 0x1e, 0x0, 0xf0, 0x7, 0x80, 0x3c, 0x1, 0xe0, 0xf, 0x0, 0x78, 0x3, 0xc0, 0x1e, 0x0, 0xf0, 0x7, 0x80, 0x3c, 0x1, 0xe0, 0xf, 0x80, 0xee, 0xe, 0x3f, 0xe0, 0x7c, 0x0}},
		/* V */ tinyfont.Glyph{Rune: 86, Width: 0xf, Height: 0x12, XAdvance: 0xf, XOffset: 0, YOffset: -17, Bitmaps: []uint8{0x60, 0x6, 0xc0, 0x1d, 0xc0, 0x31, 0x80, 0x63, 0x1, 0xc7, 0x3, 0x6, 0x6, 0xc, 0x1c, 0x1c, 0x30, 0x18, 0x60, 0x31, 0xc0, 0x73, 0x0, 0x66, 0x0, 0xdc, 0x1, 0xf0, 0x1, 0xe0, 0x3, 0xc0, 0x7, 0x0}},
		/* W */ tinyfont.Glyph{Rune: 87, Width: 0x16, Height: 0x12, XAdvance: 0x16, XOffset: 0, YOffset: -17, Bitmaps: []uint8{0xe0, 0x30, 0x1d, 0x80, 0xe0, 0x76, 0x7, 0x81, 0xd8, 0x1e, 0x6, 0x70, 0x7c, 0x18, 0xc1, 0xb0, 0xe3, 0xc, 0xc3, 0x8c, 0x33, 0xc, 0x38, 0xc6, 0x30, 0x67, 0x18, 0xc1, 0x98, 0x67, 0x6, 0x61, 0xd8, 0x1d, 0x83, 0x60, 0x3c, 0xd, 0x80, 0xf0, 0x3e, 0x3, 0xc0, 0x70, 0xf, 0x1, 0xc0, 0x18, 0x7, 0x0}},
		/* X */ tinyfont.Glyph{Rune: 88, Width: 0xf, Height: 0x12, XAdvance: 0x10, XOffset: 0, YOffset: -17, Bitmaps: []uint8{0x70, 0xe, 0x60, 0x38, 0xe0, 0x60, 0xe1, 0xc0, 0xc3, 0x1, 0xcc, 0x1, 0xf8, 0x1, 0xe0, 0x3, 0x80, 0x7, 0x80, 0x1f, 0x0, 0x33, 0x0, 0xe7, 0x3, 0x86, 0x6, 0xe, 0x1c, 0xe, 0x70, 0xc, 0xc0, 0x1c}},
		/* Y */ tinyfont.Glyph{Rune: 89, Width: 0x10, Height: 0x12, XAdvance: 0x10, XOffset: 0, YOffset: -17, Bitmaps: []uint8{0x60, 0x6, 0x70, 0xe, 0x30, 0x1c, 0x38, 0x18, 0x1c, 0x38, 0xc, 0x30, 0xe, 0x70, 0x6, 0x60, 0x3, 0xc0, 0x3, 0xc0, 0x1, 0x80, 0x1, 0x80, 0x1, 0x80, 0x1, 0x80, 0x1, 0x80, 0x1, 0x80, 0x1, 0x80, 0x1, 0x80}},
		/* Z */ tinyfont.Glyph{Rune: 90, Width: 0xd, Height: 0x12, XAdvance: 0xf, XOffset: 1, YOffset: -17, Bitmaps: []uint8{0xff, 0xff, 0xff, 0xc0, 0xe, 0x0, 0xe0, 0xe, 0x0, 0x60, 0x7, 0x0, 0x70, 0x7, 0x0, 0x30, 0x3, 0x80, 0x38, 0x3, 0x80, 0x18, 0x1, 0xc0, 0x1c, 0x0, 0xff, 0xff, 0xff, 0xc0}},
		/* [ */ tinyfont.Glyph{Rune: 91, Width: 0x4, Height: 0x17, XAdvance: 0x7, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0xff, 0xcc, 0xcc, 0xcc, 0xcc, 0xcc, 0xcc, 0xcc, 0xcc, 0xcc, 0xcf, 0xf0}},
		/* \ */ tinyfont.Glyph{Rune: 92, Width: 0x7, Height: 0x12, XAdvance: 0x7, XOffset: 0, YOffset: -17, Bitmaps: []uint8{0x81, 0x81, 0x2, 0x6, 0x4, 0x8, 0x18, 0x10, 0x20, 0x60, 0x40, 0x81, 0x81, 0x2, 0x6, 0x4}},
		/* ] */ tinyfont.Glyph{Rune: 93, Width: 0x4, Height: 0x17, XAdvance: 0x7, XOffset: 1, YOffset: -17, Bitmaps: []uint8{0xff, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x3f, 0xf0}},
		/* ^ */ tinyfont.Glyph{Rune: 94, Width: 0x9, Height: 0x9, XAdvance: 0xb, XOffset: 1, YOffset: -16, Bitmaps: []uint8{0xc, 0xe, 0x5, 0x86, 0xc3, 0x21, 0x19, 0x8c, 0x83, 0xc1, 0x80}},
		/* _ */ tinyfont.Glyph{Rune: 95, Width: 0xf, Height: 0x1, XAdvance: 0xd, XOffset: -1, YOffset: 4, Bitmaps: []uint8{0xff, 0xfe}},
		/* ` */ tinyfont.Glyph{Rune: 96, Width: 0x5, Height: 0x4, XAdvance: 0x6, XOffset: 1, YOffset: -17, Bitmaps: []uint8{0xe3, 0x8c, 0x30}},
		/* a */ tinyfont.Glyph{Rune: 97, Width: 0xc, Height: 0xd, XAdvance: 0xd, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x3f, 0x7, 0xf8, 0xe1, 0xcc, 0xc, 0x0, 0xc0, 0x1c, 0x3f, 0xcf, 0x8c, 0xc0, 0xcc, 0xc, 0xe3, 0xc7, 0xef, 0x3c, 0x70}},
		/* b */ tinyfont.Glyph{Rune: 98, Width: 0xc, Height: 0x12, XAdvance: 0xd, XOffset: 1, YOffset: -17, Bitmaps: []uint8{0xc0, 0xc, 0x0, 0xc0, 0xc, 0x0, 0xc0, 0xc, 0xf8, 0xdf, 0xcf, 0xe, 0xe0, 0x7c, 0x3, 0xc0, 0x3c, 0x3, 0xc0, 0x3c, 0x3, 0xe0, 0x6f, 0xe, 0xdf, 0xcc, 0xf8}},
		/* c */ tinyfont.Glyph{Rune: 99, Width: 0xa, Height: 0xd, XAdvance: 0xc, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x1f, 0xf, 0xe7, 0x1b, 0x83, 0xc0, 0x30, 0xc, 0x3, 0x0, 0xc0, 0x38, 0x37, 0x1c, 0xfe, 0x1f, 0x0}},
		/* d */ tinyfont.Glyph{Rune: 100, Width: 0xb, Height: 0x12, XAdvance: 0xd, XOffset: 1, YOffset: -17, Bitmaps: []uint8{0x0, 0x60, 0xc, 0x1, 0x80, 0x30, 0x6, 0x3c, 0xcf, 0xfb, 0x8f, 0xe0, 0xf8, 0xf, 0x1, 0xe0, 0x3c, 0x7, 0x80, 0xf8, 0x3b, 0x8f, 0x3f, 0x63, 0xcc}},
		/* e */ tinyfont.Glyph{Rune: 101, Width: 0xb, Height: 0xd, XAdvance: 0xd, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x1f, 0x7, 0xf1, 0xc7, 0x70, 0x3c, 0x7, 0xff, 0xff, 0xfe, 0x0, 0xc0, 0x1c, 0xd, 0xc3, 0x1f, 0xe1, 0xf0}},
		/* f */ tinyfont.Glyph{Rune: 102, Width: 0x5, Height: 0x12, XAdvance: 0x7, XOffset: 1, YOffset: -17, Bitmaps: []uint8{0x3b, 0xd8, 0xc6, 0x7f, 0xec, 0x63, 0x18, 0xc6, 0x31, 0x8c, 0x63, 0x0}},
		/* g */ tinyfont.Glyph{Rune: 103, Width: 0xb, Height: 0x12, XAdvance: 0xd, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x1e, 0x67, 0xfd, 0xc7, 0xf0, 0x7c, 0x7, 0x80, 0xf0, 0x1e, 0x3, 0xc0, 0x7c, 0x1d, 0xc7, 0x9f, 0xb1, 0xe6, 0x0, 0xc0, 0x3e, 0xe, 0x7f, 0xc7, 0xe0}},
		/* h */ tinyfont.Glyph{Rune: 104, Width: 0xa, Height: 0x12, XAdvance: 0xd, XOffset: 1, YOffset: -17, Bitmaps: []uint8{0xc0, 0x30, 0xc, 0x3, 0x0, 0xc0, 0x33, 0xcd, 0xfb, 0xc7, 0xe0, 0xf0, 0x3c, 0xf, 0x3, 0xc0, 0xf0, 0x3c, 0xf, 0x3, 0xc0, 0xf0, 0x30}},
		/* i */ tinyfont.Glyph{Rune: 105, Width: 0x2, Height: 0x12, XAdvance: 0x5, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0xf0, 0x3f, 0xff, 0xff, 0xf0}},
		/* j */ tinyfont.Glyph{Rune: 106, Width: 0x4, Height: 0x17, XAdvance: 0x6, XOffset: 0, YOffset: -17, Bitmaps: []uint8{0x33, 0x0, 0x3, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x3f, 0xe0}},
		/* k */ tinyfont.Glyph{Rune: 107, Width: 0xb, Height: 0x12, XAdvance: 0xc, XOffset: 1, YOffset: -17, Bitmaps: []uint8{0xc0, 0x18, 0x3, 0x0, 0x60, 0xc, 0x1, 0x83, 0x30, 0xc6, 0x30, 0xcc, 0x1b, 0x83, 0xf0, 0x77, 0xc, 0x61, 0x8e, 0x30, 0xe6, 0xc, 0xc1, 0xd8, 0x18}},
		/* l */ tinyfont.Glyph{Rune: 108, Width: 0x2, Height: 0x12, XAdvance: 0x5, XOffset: 1, YOffset: -17, Bitmaps: []uint8{0xff, 0xff, 0xff, 0xff, 0xf0}},
		/* m */ tinyfont.Glyph{Rune: 109, Width: 0x11, Height: 0xd, XAdvance: 0x13, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0xcf, 0x1f, 0x6f, 0xdf, 0xfc, 0x78, 0xfc, 0x18, 0x3c, 0xc, 0x1e, 0x6, 0xf, 0x3, 0x7, 0x81, 0x83, 0xc0, 0xc1, 0xe0, 0x60, 0xf0, 0x30, 0x78, 0x18, 0x3c, 0xc, 0x18}},
		/* n */ tinyfont.Glyph{Rune: 110, Width: 0xa, Height: 0xd, XAdvance: 0xd, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0xcf, 0x37, 0xef, 0x1f, 0x83, 0xc0, 0xf0, 0x3c, 0xf, 0x3, 0xc0, 0xf0, 0x3c, 0xf, 0x3, 0xc0, 0xc0}},
		/* o */ tinyfont.Glyph{Rune: 111, Width: 0xb, Height: 0xd, XAdvance: 0xd, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x1f, 0x7, 0xf1, 0xc7, 0x70, 0x7c, 0x7, 0x80, 0xf0, 0x1e, 0x3, 0xc0, 0x7c, 0x1d, 0xc7, 0x1f, 0xc1, 0xf0}},
		/* p */ tinyfont.Glyph{Rune: 112, Width: 0xc, Height: 0x11, XAdvance: 0xd, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0xcf, 0x8d, 0xfc, 0xf0, 0xee, 0x6, 0xc0, 0x3c, 0x3, 0xc0, 0x3c, 0x3, 0xc0, 0x3e, 0x7, 0xf0, 0xef, 0xfc, 0xcf, 0x8c, 0x0, 0xc0, 0xc, 0x0, 0xc0, 0x0}},
		/* q */ tinyfont.Glyph{Rune: 113, Width: 0xb, Height: 0x11, XAdvance: 0xd, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x1e, 0x67, 0xfd, 0xc7, 0xf0, 0x7c, 0x7, 0x80, 0xf0, 0x1e, 0x3, 0xc0, 0x7c, 0x1d, 0xc7, 0x9f, 0xf1, 0xe6, 0x0, 0xc0, 0x18, 0x3, 0x0, 0x60}},
		/* r */ tinyfont.Glyph{Rune: 114, Width: 0x6, Height: 0xd, XAdvance: 0x8, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0xcf, 0x7f, 0x38, 0xc3, 0xc, 0x30, 0xc3, 0xc, 0x30, 0xc0}},
		/* s */ tinyfont.Glyph{Rune: 115, Width: 0xa, Height: 0xd, XAdvance: 0xc, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x3e, 0x1f, 0xee, 0x1b, 0x0, 0xc0, 0x3c, 0x7, 0xf0, 0x3e, 0x1, 0xf0, 0x3e, 0x1d, 0xfe, 0x3e, 0x0}},
		/* t */ tinyfont.Glyph{Rune: 116, Width: 0x5, Height: 0x10, XAdvance: 0x7, XOffset: 1, YOffset: -15, Bitmaps: []uint8{0x63, 0x19, 0xff, 0xb1, 0x8c, 0x63, 0x18, 0xc6, 0x31, 0xe7}},
		/* u */ tinyfont.Glyph{Rune: 117, Width: 0xa, Height: 0xd, XAdvance: 0xd, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0xc0, 0xf0, 0x3c, 0xf, 0x3, 0xc0, 0xf0, 0x3c, 0xf, 0x3, 0xc0, 0xf0, 0x7e, 0x3d, 0xfb, 0x3c, 0xc0}},
		/* v */ tinyfont.Glyph{Rune: 118, Width: 0xc, Height: 0xd, XAdvance: 0xc, XOffset: 0, YOffset: -12, Bitmaps: []uint8{0xe0, 0x66, 0x6, 0x60, 0x67, 0xc, 0x30, 0xc3, 0xc, 0x39, 0x81, 0x98, 0x19, 0x81, 0xf0, 0xf, 0x0, 0xe0, 0xe, 0x0}},
		/* w */ tinyfont.Glyph{Rune: 119, Width: 0x11, Height: 0xd, XAdvance: 0x11, XOffset: 0, YOffset: -12, Bitmaps: []uint8{0xc1, 0xc1, 0xb0, 0xe1, 0xd8, 0x70, 0xcc, 0x2c, 0x66, 0x36, 0x31, 0x9b, 0x18, 0xcd, 0x98, 0x64, 0x6c, 0x16, 0x36, 0xf, 0x1a, 0x7, 0x8f, 0x3, 0x83, 0x80, 0xc1, 0xc0}},
		/* x */ tinyfont.Glyph{Rune: 120, Width: 0xb, Height: 0xd, XAdvance: 0xb, XOffset: 0, YOffset: -12, Bitmaps: []uint8{0x60, 0xee, 0x18, 0xc6, 0xc, 0xc1, 0xf0, 0x1c, 0x1, 0x80, 0x78, 0x1b, 0x3, 0x30, 0xc7, 0x30, 0x66, 0x6}},
		/* y */ tinyfont.Glyph{Rune: 121, Width: 0xb, Height: 0x12, XAdvance: 0xb, XOffset: 0, YOffset: -12, Bitmaps: []uint8{0xe0, 0x6c, 0xd, 0x83, 0x38, 0x63, 0xc, 0x63, 0xe, 0x60, 0xcc, 0x1b, 0x3, 0x60, 0x3c, 0x7, 0x0, 0xe0, 0x18, 0x3, 0x0, 0xe0, 0x78, 0xe, 0x0}},
		/* z */ tinyfont.Glyph{Rune: 122, Width: 0xa, Height: 0xd, XAdvance: 0xc, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0xff, 0xff, 0xf0, 0x18, 0xc, 0x7, 0x3, 0x81, 0xc0, 0x60, 0x30, 0x18, 0xe, 0x3, 0xff, 0xff, 0xc0}},
		/* { */ tinyfont.Glyph{Rune: 123, Width: 0x5, Height: 0x17, XAdvance: 0x8, XOffset: 1, YOffset: -17, Bitmaps: []uint8{0x19, 0xcc, 0x63, 0x18, 0xc6, 0x31, 0x99, 0x86, 0x18, 0xc6, 0x31, 0x8c, 0x63, 0x1c, 0x60}},
		/* | */ tinyfont.Glyph{Rune: 124, Width: 0x2, Height: 0x17, XAdvance: 0x6, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0xff, 0xff, 0xff, 0xff, 0xff, 0xfc}},
		/* } */ tinyfont.Glyph{Rune: 125, Width: 0x5, Height: 0x17, XAdvance: 0x8, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0xc7, 0x18, 0xc6, 0x31, 0x8c, 0x63, 0xc, 0x33, 0x31, 0x8c, 0x63, 0x18, 0xc6, 0x73, 0x0}},
		/* ~ */ tinyfont.Glyph{Rune: 126, Width: 0xa, Height: 0x5, XAdvance: 0xc, XOffset: 1, YOffset: -10, Bitmaps: []uint8{0x70, 0x3e, 0x9, 0xe4, 0x1f, 0x3, 0x80}},
	},

	YAdvance: 0x1d,
}
