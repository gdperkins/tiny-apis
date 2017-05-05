package web

import (
	"regexp"

	"github.com/gdperkins/tiny-apis/models"
	"github.com/gdperkins/tiny-apis/services"

	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

//Color APIs belong in this go file, validation etc. Any logic should be in a service

var (
	hexRegEx = "^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})"
	rgbRegEx = "^\\((\\d{1,3}),(\\d{1,3}),(\\d{1,3})\\)$"
)

// ConvertColor takes a color input in one of the following
// formats: RGB, HEX, CMYK or HSV and returns a color cummary
// struct with all the missing formats filled out
func ConvertColor(c *gin.Context) {

	hex := c.Query("hex")
	if hex != "" {
		match, _ := regexp.Match(hexRegEx, []byte(hex))
		if match == false {
			c.JSON(http.StatusBadRequest, models.NewError("Invalid request", 2,
				"Invalid hex format."))
			return
		}
		c.JSON(http.StatusOK, services.ColorSummaryFromHex(hex))
	}

	rgb := c.Query("rgb")
	if rgb != "" {
		match, _ := regexp.Match(rgbRegEx, []byte(rgb))
		if match == false {
			c.JSON(http.StatusBadRequest, models.NewError("Invalid request", 2,
				"Invalid rgb format."))
			return
		}
		c.JSON(http.StatusOK, services.ColorSummaryFromRgb(rgb))
	}

	hsv := c.Query("hsv")
	if hsv != "" {

	}

	cmyk := c.Query("cmyk")
	if cmyk != "" {

	}

	if hex == "" && rgb == "" && hsv == "" && cmyk == "" {
		c.JSON(http.StatusBadRequest, models.NewError("Invalid request", 1,
			"Missing color code parameter. Hex, hsv, rgb or cmyk required."))
		return
	}
}
