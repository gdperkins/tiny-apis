package services

import (
	"encoding/hex"
	"fmt"
	"math"
)

// ColorSummary represents a colour in multiple
// standard colour code formats
type ColorSummary struct {
	RGB  rgb    `json:"rgb"`
	HEX  string `json:"hex"`
	CMYK cmyk   `json:"cmyk"`
	HSV  hsv    `json:"hsv"`
}

// color is the RGB representation of
// the value the user colour input
type rgb struct {
	Red    int    `json:"red"`
	Green  int    `json:"green"`
	Blue   int    `json:"blue"`
	Output string `json:"output"`
}

type cmyk struct {
	Cyan    float64
	Magenta float64
	Yellow  float64
	Black   float64
	Output  string
}

type hsv struct {
	Hue        string
	Saturation string
	Value      string
	Output     string
}

// ColorSummaryFromHex receives a hex input and generates the
// rest of the summary (CMYK, HSV, RGB)
func ColorSummaryFromHex(input string) ColorSummary {
	decoded, _ := hex.DecodeString(input[1:len(input)])
	fmt.Println(decoded)

	rgb := rgb{
		Red:   int(decoded[0]),
		Green: int(decoded[1]),
		Blue:  int(decoded[2]),
		Output: fmt.Sprintf("(%d,%d,%d)",
			int(decoded[0]), int(decoded[1]), int(decoded[2])),
	}

	return ColorSummary{
		RGB:  rgb,
		HEX:  input,
		CMYK: cmykFromRgb(&rgb),
		HSV:  hsvFromRgb(&rgb),
	}
}

func cmykFromRgb(rgb *rgb) cmyk {
	if rgb.Red == 0 && rgb.Green == 0 && rgb.Blue == 0 {
		return cmyk{0, 0, 0, 1, "0, 0, 0, 1"}
	}

	c := float64(1 - (rgb.Red / 255))
	m := float64(1 - (rgb.Green / 255))
	y := float64(1 - (rgb.Blue / 255))
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
