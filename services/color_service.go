package services

import (
	"encoding/hex"
	"fmt"
	"math"
)

// ColorSummary represents a colour in multiple
// standard colour code formats
type ColorSummary struct {
	HEX  string `json:"hex"`
	RGB  rgb    `json:"rgb"`
	CMYK cmyk   `json:"cmyk"`
	HSV  hsv    `json:"hsv"`
}

// color is the RGB representation of
// the value the user colour input
type rgb struct {
	R      uint8  `json:"r"`
	G      uint8  `json:"g"`
	B      uint8  `json:"b"`
	Output string `json:"output"`
}

type cmyk struct {
	C      uint8  `json:"c"`
	M      uint8  `json:"m"`
	Y      uint8  `json:"y"`
	K      uint8  `json:"k"`
	Output string `json:"output"`
}

type hsv struct {
	H      string `json:"h"`
	S      string `json:"s"`
	V      string `json:"v"`
	Output string `json:"output"`
}

// ColorSummaryFromHex receives a hex input and generates the
// rest of the summary (CMYK, HSV, RGB)
func ColorSummaryFromHex(input string) ColorSummary {
	rgb := rgbFromHex(input)

	cmyk := cmykFromRgb(&rgb)

	return ColorSummary{
		RGB:  rgb,
		HEX:  input,
		CMYK: cmyk,
		HSV:  hsvFromRgb(&rgb),
	}
}

func rgbFromHex(input string) rgb {
	decoded, _ := hex.DecodeString(input[1:len(input)])

	return rgb{
		uint8(decoded[0]),
		uint8(decoded[1]),
		uint8(decoded[2]),
		fmt.Sprintf("(%d,%d,%d)",
			int(decoded[0]), int(decoded[1]), int(decoded[2])),
	}
}

func cmykFromRgb(rgb *rgb) cmyk {
	if rgb.R == 0 && rgb.G == 0 && rgb.B == 0 {
		return cmyk{0, 0, 0, 1, "0, 0, 0, 1"}
	}

	r := 1 - (float64(rgb.R) / float64(255))
	g := 1 - (float64(rgb.G) / float64(255))
	b := 1 - (float64(rgb.B) / float64(255))
	bl := min(r, min(g, b))

	c := round((r - bl) / (1 - bl) * 100)
	m := round((g - bl) / (1 - bl) * 100)
	y := round((b - bl) / (1 - bl) * 100)
	k := round(bl * 100)

	return cmyk{c, m, y, k,
		fmt.Sprintf("(%d, %d, %d, %d)", c, m, y, k),
	}
}

func hsvFromRgb(rgb *rgb) hsv {

	return hsv{}
}

func min(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

func round(val float64) uint8 {
	var round float64
	pow := math.Pow(10, float64(0))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= 0.5 {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	return uint8(round / pow)
}
