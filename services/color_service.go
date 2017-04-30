package services

import (
	"encoding/hex"
	"fmt"
)

// ColorSummary represents a colour in multiple
// standard colour code formats
type ColorSummary struct {
	RGB  rgb                    `json:"rgb"`
	HEX  string                 `json:"hex"`
	CMYK map[string]interface{} `json:"cmyk"`
	HSV  string                 `json:"hsv"`
}

// color is the RGB representation of
// the value the user colour input
type rgb struct {
	R     int    `json:"r"`
	G     int    `json:"g"`
	B     int    `json:"b"`
	Value string `json:"value"`
}

type cmyk struct {
}

type hsv struct {
}

var (
	// the currentRgb for converting to other color formats
	currentRgb rgb
)

// GenerateColorSummary converts the input to other color formats
// Hex, RGB, CMYK and HSV.
func GenerateColorSummary(input string) ColorSummary {

	//if hex
	decoded, _ := hex.DecodeString(input[1:len(input)])
	fmt.Println(decoded)

	currentRgb = rgb{
		R: int(decoded[0]),
		G: int(decoded[1]),
		B: int(decoded[2]),
		Value: fmt.Sprintf("(%d,%d,%d)",
			int(decoded[0]), int(decoded[1]), int(decoded[2])),
	}

	//if cmyk

	//if hsv

	return ColorSummary{
		RGB: currentRgb,
		HEX: input,
	}
}

func getRgb() (r, g, b int) {
	return currentRgb.R, currentRgb.G, currentRgb.B
}

func getHex() string {
	return ""
}

func getCmyk() (c, m, y, k float32) {
	return 0, 0, 0, 0
}

func getHsv() (h, s, v float32) {
	return 0, 0, 0
}
