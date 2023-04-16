package controllers

import (
	"fmt"
	"movies/database"
	"movies/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {

	// get the email/pass off req body

	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body ",
		})

		return

	}

	// hash the password

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to hash the password ",
		})

		return
	}

	// create the user

	user := models.User{Email: body.Email, Password: string(hash)}
	database.StartDb()
	var errors error

	errors = database.CreateUser(user)

	if errors != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create user ",
		})

		return
    }

	database.CloseDb()


	//respond
	c.JSON(http.StatusOK, gin.H{})

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
	database.StartDb()

	result , err := database.GetUserFromEmail(body.Email)


	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "invalid email or password",
		})

		return
    }


	for result.Next(){
		if err := result.Scan(&user.ID , &user.Email , &user.Password ); err != nil {
            panic(err)
        }
	}

	if err := result.Err(); err != nil {
        panic(err)
    }

	database.CloseDb()
	



	// compare the sent in pass with saved hash pass

	err = bcrypt.CompareHashAndPassword([]byte(user.Password) , []byte(body.Password))

	if err != nil{

		c.JSON(http.StatusBadRequest , gin.H{
			"error" : "invalid email or password",
		})

		return
	}

	//generate a jwt token 

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

	// send the jwt  and user back



	fmt.Println(tokenString)

	var response models.Response
	response.User = user 
	response.Jwt = tokenString


	c.JSON(http.StatusOK ,  response )

}

func Follow(c *gin.Context){


	// getting the body of request from fronend

	var body struct{
		UserEmail string
		FollowedUserName string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body ",
		})
		return
	}

	// first querry to find the first user
	var user models.User 
	database.StartDb()

	result , err := database.GetUserFromEmail(body.UserEmail)


	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "invalid email",
		})

		return
    }


	for result.Next(){
		if err := result.Scan(&user.ID , &user.Email , &user.Password ); err != nil {
            panic(err)
        }
	}

	if err := result.Err(); err != nil {
        panic(err)
    }

	// second querry to find the person should be followd 
	// by finding it from username


	result , err = database.GetUserFromUserName(body.FollowedUserName)

	var secondUser models.User 

	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "invalid username",
		})

		return
    }


	for result.Next(){
		if err := result.Scan(&secondUser.ID); err != nil {
            panic(err)
        }
	}

	if err := result.Err(); err != nil {
        panic(err)
    }

	// querry to insert the people involve in follow 

	err = database.InserFollow(user.ID , secondUser.ID)


	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create follow entity ",
		})

		return
    }

	database.CloseDb()

}
