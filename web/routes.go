package web

import (
	"net/http"

	"github.com/gdperkins/tiny-apis/apis"

	"fmt"

	"gopkg.in/gin-gonic/gin.v1"
)

// RegisterRoutes adds all the APIs to be served
func RegisterRoutes(s *Server) {
	s.GET("/test", func(c *gin.Context) {
		message := ""
		name := c.Query("name")
		if name != "" {
			message = fmt.Sprintf("Hello %s, lets make some APIs.", name)
		} else {
			message = "Hello guest, lets make some APIs."
		}
		c.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	})

	clrs := s.Group("api/v1/colors")
	{
		clrs.GET("/convert", apis.ConvertWebColour)
	}
}
