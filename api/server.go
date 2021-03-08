package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/omaciel/gobutsu/db/sqlc"
	"github.com/omaciel/gobutsu/util"
)

// Server serves HTTP requests
type Server struct {
	config util.Config
	app    db.App
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config util.Config, app db.App) *Server {
	server := &Server{
		config: config,
		app:    app,
	}
	router := gin.Default()

	router.POST("/testcases", server.createTestCase)
	router.GET("/testcases/:id", server.getTestCase)
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
