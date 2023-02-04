package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	/// group of urls

	sth := gin.Default()

	var firstGroup = sth.Group("/shop")

	firstGroup.GET("/hello", func(c *gin.Context) {

		c.String(200, "hello world")
	})

	var secondGroup = sth.Group("/dance")

	secondGroup.GET("/hello", func(c *gin.Context) {

		c.String(200, "hello world")
	})

	sth.Run(":8080")

}
