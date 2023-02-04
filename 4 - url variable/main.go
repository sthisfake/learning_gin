package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	sth := gin.Default()

	/// c.param

	sth.GET("/hello/:id", func(c *gin.Context) {
		var id = c.Param("id")
		fmt.Println("id = ", id)
		c.String(200, "hello "+id)
	})

	/// c.query

	sth.GET("/sth/:id", func(c *gin.Context) {
		var id = c.Query("id")
		c.String(200, "hello "+id)
	})

	/// c.defaultquery

	sth.GET("/how/:id", func(c *gin.Context) {
		var id = c.DefaultQuery("id", "500")
		c.String(200, "hello "+id)
	})

	/// c.postform

	sth.POST("/test/:id", func(c *gin.Context) {
		var id = c.PostForm("id")
		c.String(200, "hello "+id)
	})

	/// c.defaultpostform

	sth.POST("/pou/:id", func(c *gin.Context) {
		var id = c.DefaultPostForm("id", "100")
		c.String(200, "hello "+id)
	})

	/// c.getheader

	sth.POST("/lolo/:id", func(c *gin.Context) {
		var id = c.GetHeader("id")
		c.String(200, "hello "+id)
	})

	sth.GET("/hello", func(c *gin.Context) {

		c.String(200, "hello world")
	})

	sth.Run(":8080")

}
