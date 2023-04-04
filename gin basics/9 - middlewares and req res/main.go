package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	UserName string
	Password string
}

func main() {

	server := gin.Default()

	// server.Use(func(ctx *gin.Context) {
	// 	var requestID = ctx.GetHeader("X-Request-Id")

	// 	if len(requestID) == 0 {
	// 		var id = uuid.New().String()

	// 		ctx.Writer.Header().Add("X-Request-Id", id)

	// 	} else {
	// 		ctx.Writer.Header().Add("X-Request-Id", requestID)
	// 	}
	// 	fmt.Println("loadinggg.....")
	// })

	var accounts = map[string]string{
		"pouya": "tey",
		"kir":   "khar",
	}

	var sth = gin.BasicAuth(accounts)

	server.GET("/test", sth, func(ctx *gin.Context) {
		ctx.String(200, "sth")
	})

	server.GET("/hello", func(ctx *gin.Context) {
		ctx.String(200, "helooooooo")
	})

	server.GET("/res", func(ctx *gin.Context) {

		var user = User{
			UserName: "someone",
			Password: "1245",
		}

		ctx.JSON(200, user)
	})

	server.Run(":8080")

}
