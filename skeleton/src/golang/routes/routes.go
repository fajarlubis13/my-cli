package routes

import (
	"{{ toDelimeted .ProjectName 45 }}/handlers"

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
		v2.POST("/{{ toDelimeted .ProjectName 45 }}", handler.CreateOne)
		v2.PUT("/{{ toDelimeted .ProjectName 45 }}/:id", handler.UpdateOneByID)
		v2.GET("/{{ toDelimeted .ProjectName 45 }}/:id", handler.GetOneByID)
		v2.DELETE("/{{ toDelimeted .ProjectName 45 }}/:id", handler.DeleteOneByID)
		v2.GET("/{{ toDelimeted .ProjectName 45 }}", handler.GetAll)
	}

	return r
}
