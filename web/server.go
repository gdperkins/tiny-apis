package web

import (
	"gopkg.in/gin-gonic/gin.v1"
)

// Server is the base struct for hosting using the gin Engine
// embedded type as the inner workings
type Server struct {
	*gin.Engine
}

// NewServer creates a new Server struct to host the application.
// Any middleware or environment configuration settings will be
// applied here.
func NewServer() *Server {

	s := Server{
		gin.New(),
	}

	// add middleware

	s.Use(gin.Logger())
	s.Use(gin.Recovery())

	// add all the api routes
	RegisterRoutes(&s)

	return &s
}
