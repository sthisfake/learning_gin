package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/// example 1

// func main() {
// 	r := gin.Default()
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "pong",
// 		})
// 	})
// 	r.GET("/hey", func(nothing *gin.Context) {
// 		nothing.JSON(http.StatusOK, gin.H{
// 			"kir": "khar",
// 		})
// 	})
// 	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
// }

/// end example 1

//**************************************************************************//

/// example 2 : list of http methods

// func main() {
// // Creates a gin router with default middleware:
// // logger and recovery (crash-free) middleware
// router := gin.Default()

// router.GET("/someGet", getting)
// router.POST("/somePost", posting)
// router.PUT("/somePut", putting)
// router.DELETE("/someDelete", deleting)
// router.PATCH("/somePatch", patching)
// router.HEAD("/someHead", head)
// router.OPTIONS("/someOptions", options)

// // By default it serves on :8080 unless a
// // PORT environment variable was defined.
// router.Run()
// // router.Run(":3000") for a hard coded port
// }

/// example 2 end

//**************************************************************************//

/// example 3

// func main() {
// 	router := gin.Default()

// 	// This handler will match /user/john but will not match /user/ or /user
// 	router.GET("/user/:name", func(c *gin.Context) {
// 		name := c.Param("name")
// 		c.String(http.StatusOK, "Hello %s", name)
// 	})

// 	// However, this one will match /user/john/ and also /user/john/send
// 	// If no other routers match /user/john, it will redirect to /user/john/
// 	router.GET("/user/:name/*action", func(c *gin.Context) {
// 		name := c.Param("name")
// 		action := c.Param("action")
// 		message := name + " is " + action
// 		c.String(http.StatusOK, message)
// 	})

// 	// For each matched request Context will hold the route definition
// 	router.POST("/user/:name/*action", func(c *gin.Context) {
// 		b := c.FullPath() == "/user/:name/*action" // true
// 		c.String(http.StatusOK, "%t", b)
// 	})

// 	// This handler will add a new router for /user/groups.
// 	// Exact routes are resolved before param routes, regardless of the order they were defined.
// 	// Routes starting with /user/groups are never interpreted as /user/:name/... routes
// 	router.GET("/user/groups", func(c *gin.Context) {
// 		c.String(http.StatusOK, "The available groups are [...]")
// 	})

// 	router.Run(":8080")
// }

/// example 3 end

//**************************************************************************//

/// example 4

// func main() {
// 	router := gin.Default()

// 	// Query string parameters are parsed using the existing underlying request object.
// 	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
// 	router.GET("/welcome", func(c *gin.Context) {
// 		firstname := c.DefaultQuery("firstname", "Guest")
// 		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

// 		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
// 	})
// 	router.Run(":8080")
// }

func main() {
	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	router.Run(":8080")
}
