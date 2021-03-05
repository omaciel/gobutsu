package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/omaciel/gobutsu/db/sqlc"
)

// Server serves HTTP requests
type Server struct {
	app *db.App
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(app *db.App) *Server {
	server := &Server{
		app:    app,
	}
	router := gin.Default()

	router.POST("/testcases", server.createTestCase)
	router.GET("/testcases/:testcaseid", server.getTestCase)
	router.GET("/testcases/", server.listTestCase)

	server.router = router
	return server
}

// Start runs the HTTP server on the specified address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}