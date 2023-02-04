package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Moshtari struct {
	Email         string `json:"email" binding:"required,email"`
	Password      string `json:"password" binding:"required,password"`
	Role          string `json:"role" binding:"required,oneof=STG STH"`
	StreetAddress string `json:"streetAddress"`
	StreetNumber  int    `json:"streetNumber" binding:"required_with=StreetAddress"`
}

func verifyPassword(f1 validator.FieldLevel) bool {
	var regex = regexp.MustCompile("\\w{8,}")
	var password = f1.Field().String()
	return regex.MatchString(password)
}

func main() {

	sth := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("password", verifyPassword)
	}

	sth.POST("/moshtari", func(ctx *gin.Context) {
		var cust Moshtari

		if e := ctx.ShouldBindJSON(&cust); e != nil {
			ctx.String(http.StatusBadRequest, e.Error())
			return
		}

		fmt.Println(cust)
		ctx.String(200, "hello  "+cust.Email+" "+cust.Role+" "+cust.StreetAddress+" "+strconv.Itoa(cust.StreetNumber)+" \n "+cust.Password)
	})

	sth.Run(":8080")

}
