package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	sth := gin.Default()

	sth.GET("/hello", func(c *gin.Context) {

		c.String(200, "hello world")

	})

	sth.Run(":8080")

}
