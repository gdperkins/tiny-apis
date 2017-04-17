package web

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

// RegisterRoutes adds all the APIs to be served
func RegisterRoutes(s *Server) {
	s.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello %s, lets make some APIs.", c.Query("name"))
	})
}
