package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	db "github.com/omaciel/gobutsu/db/sqlc"
	"github.com/omaciel/gobutsu/util"
)

func newTestServer(t *testing.T, app db.App) *Server {
	config := util.Config{}

	server := NewServer(config, app)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
