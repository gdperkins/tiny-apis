package services

import (
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// ColorSummary represents a colour in multiple
// standard colour code formats
type ColorSummary struct {
	HEX  string `json:"hex"`
	RGB  rgb    `json:"rgb"`
	CMYK cmyk   `json:"cmyk"`
	HSV  hsv    `json:"hsv"`
}

// ColorSummaryFromHex generates CMYK, HSV, RGB from the hex input.
func ColorSummaryFromHex(input string) ColorSummary {
	rgb := rgbFromHex(input)
	return ColorSummary{
		RGB:  rgb,
		HEX:  input,
		CMYK: cmykFromRgb(&rgb),
		HSV:  hsvFromRgb(&rgb),
	}
}

// ColorSummaryFromRgb generates CMYK, HSV, HEX from the Rgb input.
func ColorSummaryFromRgb(input string) ColorSummary {
	input = strings.TrimLeft(input, "(")
	input = strings.TrimRight(input, ")")

	rgbs := strings.Split(input, ",")
	r, _ := strconv.Atoi(rgbs[0])
	g, _ := strconv.Atoi(rgbs[1])
	b, _ := strconv.Atoi(rgbs[2])

	color := rgb{
		uint(r),
		uint(g),
		uint(b),
		fmt.Sprintf("(%d,%d,%d)", r, g, b),
	}

	return ColorSummary{
		RGB:  color,
		HEX:  hexFromRgb(&color),
		CMYK: cmykFromRgb(&color),
		HSV:  hsvFromRgb(&color),
	}
}

type rgb struct {
	R      uint   `json:"r"`
	G      uint   `json:"g"`
	B      uint   `json:"b"`
	Output string `json:"output"`
}

// ColorSummaryFromHsv generates CMYK, RGB, HEX from the Hsv input
func ColorSummaryFromHsv(input string) ColorSummary {
	hsvs := strings.Split(input, ",")
	h, _ := strconv.Atoi(hsvs[0])
	s, _ := strconv.Atoi(hsvs[1])
	v, _ := strconv.Atoi(hsvs[2])

	hsv := hsv{
		H: uint(h),
		S: uint(s),
		V: uint(v),
	}

	rgb := rgbFromHsv(hsv)

	return ColorSummary{
		HSV:  hsv,
		RGB:  rgb,
		CMYK: cmykFromRgb(&rgb),
		HEX:  hexFromRgb(&rgb),
	}
}

type hsv struct {
	H      uint   `json:"h"`
	S      uint   `json:"s"`
	V      uint   `json:"v"`
	Output string `json:"output"`
}

type cmyk struct {
	C      uint   `json:"c"`
	M      uint   `json:"m"`
	Y      uint   `json:"y"`
	K      uint   `json:"k"`
	Output string `json:"output"`
}

func hexFromRgb(rgb *rgb) string {
	return fmt.Sprintf("#%v%v%v",
		strconv.FormatInt(int64(rgb.R), 16),
		strconv.FormatInt(int64(rgb.G), 16),
		strconv.FormatInt(int64(rgb.B), 16),
	)
}

func rgbFromHsv(color hsv) rgb {

	return rgb{}
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
