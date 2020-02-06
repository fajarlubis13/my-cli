package routes

import (
	"hk-jadwal-teknik/handlers"

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
		v2.POST("/jadwal-teknik", handler.CreateOne)
		v2.PUT("/jadwal-teknik/:id", handler.UpdateOneByID)
		v2.GET("/jadwal-teknik/:id", handler.GetOneByID)
		v2.DELETE("/jadwal-teknik/:id", handler.DeleteOneByID)
		v2.GET("/jadwal-teknik", handler.GetAll)
	}

	return r
}
