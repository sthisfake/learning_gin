package middleware

import "github.com/gin-gonic/gin"

func RequireAuth(c *gin.Context){

	//get the cookie off req 

	tokenString , err := c.Cookie("authorization")

	// decode / validate it


	// check the exp 

	// find the user with token sub

	// attach to req

	//continue

	c.Next()

}