package routes

import (
	"hk-pengiriman/handlers"

	"github.com/gin-gonic/gin"
)

// NewRoutes ...
func NewRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	handler := handlers.NewHandlers()

	v2 := r.Group("/v2")
	{
		v2.POST("/hk-pengiriman", handler.CreateOne)
		v2.PUT("/hk-pengiriman/:id", handler.UpdateOneByID)
		v2.GET("/hk-pengiriman/:id", handler.GetOneByID)
		v2.DELETE("/hk-pengiriman/:id", handler.DeleteOneByID)
		v2.GET("/hk-pengiriman", handler.GetAll)
	}

	return r
}
