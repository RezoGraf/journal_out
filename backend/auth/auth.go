package auth

import (
	"github.com/gin-gonic/gin"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"log"
	"../models"
	"fmt"
)

type JwtClaims struct {
	Id 			string	  `json:"id"`
	Name        string    `json:"name"`
	jwt.StandardClaims
}

func Login(c *gin.Context, userInfo []*models.ModelUser) string{

	//TODO: create jwt token
	id  := userInfo[0].Id
	fam := userInfo[0].Fam
	im  := userInfo[0].Im
	ot  := userInfo[0].Ot
	token, err := createJwtToken(id, fam, im, ot)
	if err != nil {
		log.Println("Error Creating JWT token", err)
	}
	return token
}

func createJwtToken(id, fam, im, ot string) (string, error)  {
	fmt.Println("creating jwt tokin", id, fam, im, ot)
	claims := JwtClaims{
		id,
		fam +" "+ im +" "+ot,
		jwt.StandardClaims{
			Id: string(id),
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
