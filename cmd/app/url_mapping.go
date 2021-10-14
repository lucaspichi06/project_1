package app

import "github.com/gin-gonic/gin"

// mapRoutes to create the routes
func (c *controllers) mapRoutes(router *gin.Engine) {
	router.GET("/ping", c.status.HandlePing)
	router.POST("/produce", c.produce.Add)
	router.GET("/produce", c.produce.Fetch)
	router.DELETE("/produce/:id", validateId(), c.produce.Delete)
}
