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
	R      int    `json:"r"`
	G      int    `json:"g"`
	B      int    `json:"b"`
	Output string `json:"output"`
}

type cmyk struct {
	C      float64 `json:"c"`
	M      float64 `json:"m"`
	Y      float64 `json:"y"`
	K      float64 `json:"k"`
	Output string  `json:"output"`
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
	decoded, _ := hex.DecodeString(input[1:len(input)])
	fmt.Println(decoded)

	rgb := rgb{
		int(decoded[0]),
		int(decoded[1]),
		int(decoded[2]),
		fmt.Sprintf("(%d,%d,%d)",
			int(decoded[0]), int(decoded[1]), int(decoded[2])),
	}

	return ColorSummary{
		RGB:  rgb,
		HEX:  input,
		CMYK: cmyk{},
		HSV:  hsvFromRgb(&rgb),
	}
}

func cmykFromRgb(rgb *rgb) cmyk {
	if rgb.R == 0 && rgb.G == 0 && rgb.B == 0 {
		return cmyk{0, 0, 0, 1, "0, 0, 0, 1"}
	}

	c := float64(1 - (rgb.R / 255))
	m := float64(1 - (rgb.G / 255))
	y := float64(1 - (rgb.B / 255))
	min := float64(math.Min(c, math.Min(m, y)))

	return cmyk{
		(c - min) / (1 - min),
		(m - min) / (1 - min),
		(y - min) / (1 - min),
		min,
		fmt.Sprintf("%f, %f, %f, %f", c, m, y, min),
	}
}

func hsvFromRgb(rgb *rgb) hsv {

	return hsv{}
}
