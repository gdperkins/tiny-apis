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
)

// ConvertColor takes a color input in one of the following
// formats: RGB, HEX, CMYK or HSV and returns a color cummary
// struct with all the missing formats filled out
func ConvertColor(c *gin.Context) {

	//hex, cmyk, hsv or rgb can be inputted here, need to check for them in this order:
	//RGB, hex, cmyk, hsv

	h := c.Query("hex")
	if h == "" {
		c.JSON(http.StatusBadRequest, models.NewError("Invalid request", 1, "Missing hex parameter."))
		return
	}
	match, _ := regexp.Match(hexRegEx, []byte(h))
	if match == false {
		c.JSON(http.StatusBadRequest, models.NewError("Invalid request", 2, "Invalid hex format."))
		return
	}

	c.JSON(http.StatusOK, services.ColorSummaryFromHex(h))
}
