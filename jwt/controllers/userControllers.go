package controllers

import (
	"fmt"
	initializers "jwt/initializer"
	"jwt/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {

	// get the email/pass off req body 

	var body struct {
		Email string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest , gin.H{
			"error" : "failed to read body ",
		})

		return
		
	}

	// hash the password 
	
	hash , err := bcrypt.GenerateFromPassword([]byte(body.Password) , 10)

	if err != nil{

		c.JSON(http.StatusBadRequest , gin.H{
			"error" : "failed to hash the password ",
		})

		return
	}

	// create the user 

	user := models.User{Email : body.Email , Password : string(hash)}
	fmt.Println(user)
	result := initializers.DB.Create(&user)
	fmt.Println(result)

	if result.Error != nil {
		
		c.JSON(http.StatusBadRequest , gin.H{
			"error" : "failed to create user ",
		})

		return
	}


	//respond 
	c.JSON(http.StatusOK , gin.H{})


}


func Login(c *gin.Context){

	// get the eamil and password off req body

	var body struct {
		Email string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest , gin.H{
			"error" : "failed to read body ",
		})

		return
		
	}


	// look up requested user

	var user models.User 
	initializers.DB.First(&user , "email = ?" , body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest , gin.H{
			"error" : "invalid email or password",
		})

		return
	}

	// compare the sent in pass with saved hash pass

	err := bcrypt.CompareHashAndPassword([]byte(user.Password) , []byte(body.Password))

	if err != nil{

		c.JSON(http.StatusBadRequest , gin.H{
			"error" : "invalid email or password",
		})

		return
	}

	// generate a jwt token 

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour *24 *30 ).Unix(),
	})
	
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest , gin.H{
			"error" : "failed to create token",
		})

		return
	}

	// send the jwt back

	// as a cookie :
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("authorization" , tokenString , 3600 * 24 * 30 , "" , "" , false , true)

	// as a response :

	c.JSON(http.StatusOK , gin.H{
		"token" : tokenString,
	})



	
}

func Validate(c *gin.Context){

	user , _ := c.Get("user")

	c.JSON(http.StatusOK , gin.H{
		"massage" : user,
	})
}