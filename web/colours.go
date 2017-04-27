package web

import (
	"fmt"
	"regexp"

	"github.com/gdperkins/tiny-apis/models"

	"net/http"

	"strings"

	"gopkg.in/gin-gonic/gin.v1"
)

// ColourSummary represents a colour in multiple
// standard colour code formats
type ColourSummary struct {
	RGB  string `json:"rgb"`
	HEX  string `json:"hex"`
	CMYK string `json:"cmyk"`
	HSV  string `json:"hsv"`
}

var (
	hexRegEx = "^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})"
)

// ConvertWebColour takes a Hex input and converts
func ConvertWebColour(c *gin.Context) {
	hex := c.Query("code")
	if hex == "" {
		c.JSON(http.StatusBadRequest, models.NewError("Invalid request", 1, "Missing hex parameter."))
		return
	}

	match, _ := regexp.Match(hexRegEx, []byte(hex))
	if match == false {
		c.JSON(http.StatusBadRequest, models.NewError("Invalid request", 1, "Invalid hex format."))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"summary": ColourSummary{
			RGB: strings.ToUpper(fmt.Sprintf("%s,%s,%s", hex[1:3], hex[3:5], hex[5:7])),
			HEX: hex},
	})
}
