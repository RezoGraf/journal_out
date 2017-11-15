package auth

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"log"
)

type JwtClaims struct {
	Name        string    `json:"name"`
	jwt.StandardClaims
}

func Login(c *gin.Context) {
	//c.Header("Access-Control-Allow-Origin", "*")
	
	//TODO: create jwt token
	token, err := createJwtToken()
	fmt.Println(token)
	if err != nil {
		log.Println("Error Creating JWT token", err)
	}
	c.JSON(200, gin.H{
		"message": "You were logged in!",
		"token": token,
	})
}

func createJwtToken() (string, error)  {
	claims := JwtClaims{
		"jack",
		jwt.StandardClaims{
			Id: "main_user_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := rawToken.SignedString([]byte("mySecret"))
	if err != nil {
		return "", err
	}
	return token, nil
}
