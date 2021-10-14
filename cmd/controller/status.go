package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// StatusController struct provides the handler for a health check endpoint.
type status struct{}

func NewStatusController() *status {
	return &status{}
}

// HandlePing returns a successful pong answer to all HTTP requests.
func (s *status) HandlePing(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
