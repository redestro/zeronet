package handlers

import (
	"github.com/gin-gonic/gin"
)

//InitRoutes exports the router
func InitRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "You have arrived at zeronet",
		})
	})
	return r
}
