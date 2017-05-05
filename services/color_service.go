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
	R      uint   `json:"r"`
	G      uint   `json:"g"`
	B      uint   `json:"b"`
	Output string `json:"output"`
}

type cmyk struct {
	C      uint   `json:"c"`
	M      uint   `json:"m"`
	Y      uint   `json:"y"`
	K      uint   `json:"k"`
	Output string `json:"output"`
}

type hsv struct {
	H      uint   `json:"h"`
	S      uint   `json:"s"`
	V      uint   `json:"v"`
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
		uint(decoded[0]),
		uint(decoded[1]),
		uint(decoded[2]),
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
	bl := math.Min(r, math.Min(g, b))

	c := round((r - bl) / (1 - bl) * 100)
	m := round((g - bl) / (1 - bl) * 100)
	y := round((b - bl) / (1 - bl) * 100)
	k := round(bl * 100)

	return cmyk{c, m, y, k,
		fmt.Sprintf("(%d, %d, %d, %d)", c, m, y, k),
	}
}

func hsvFromRgb(rgb *rgb) hsv {
	r := (float64(rgb.R) / float64(255))
	g := (float64(rgb.G) / float64(255))
	b := (float64(rgb.B) / float64(255))

	min := math.Min(r, math.Min(g, b))
	v := math.Max(r, math.Max(g, b))

	s := 0.0
	delta := v - min
	if delta != 0.0 {
		s = delta / v
	}

	h := 0.0
	if v != min {
		switch v {
		case r:
			h = math.Mod((g-b)/delta, 6.0)
		case g:
			h = (b-r)/delta + 2.0
		case b:
			h = (r-g)/delta + 4.0
		}
		h *= 60
		if h < 0.0 {
			h += 360
		}
	}

	fh, fs, fv := round(h), round(s*100), round(v*100)

	return hsv{fh, fs, fv,
		fmt.Sprintf("(%d\u00B0, %d%%, %d%%)", fh, fs, fv)}
}

func round(val float64) uint {
	var round float64
	pow := math.Pow(10, float64(0))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= 0.5 {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	return uint(round / pow)
}
