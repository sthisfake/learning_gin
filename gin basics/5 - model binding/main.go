package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type urlBinding struct {
	ID string `uri:"id"`
}

type somethig struct {
	ID   string `JSON:"id"`
	Name string `JSON:"name"`
}

type test struct {
	ID   string `form:"id"`
	Name string `form:"name"`
}

type head struct {
	RequestId string `header:"kossher"`
}

func main() {

	sth := gin.Default()

	/// c.shouldbindurl

	sth.GET("/hello/:id", func(c *gin.Context) {
		var binding urlBinding
		c.ShouldBindUri(&binding)
		fmt.Println("binding ", binding)
		c.String(200, "hello "+binding.ID)
	})

	/// c.shouldbindJSON

	sth.POST("/someone", func(c *gin.Context) {

		var binding somethig
		c.ShouldBindJSON(&binding)
		fmt.Println("binding ", binding)
		c.String(200, "hello "+binding.ID+binding.Name)
	})

	/// c.shouldbind

	sth.POST("/testing", func(c *gin.Context) {
		var binding test
		c.ShouldBind(&binding)
		fmt.Println("binding ", binding)
		c.String(200, "hello "+binding.ID+binding.Name)
	})

	/// c.shouldbindheader

	sth.POST("/heading", func(c *gin.Context) {
		var binding head
		c.ShouldBindHeader(&binding)
		fmt.Println("binding ", binding)
		c.String(200, "hello "+binding.RequestId)
	})

	sth.Run(":8080")

}
