package controllers

import (
	"github.com/gin-gonic/gin"
	"../models"
	"../auth"
	"fmt"
	"net/http"
	"github.com/dgrijalva/jwt-go"
)

func FindPatientInArena(c *gin.Context)  {

	tokenString := c.GetHeader("Authorization")
	fmt.Println("!", tokenString)



	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("mySecret"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["id"], claims["name"])
	} else {
		fmt.Println(err)
	}

	//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	//c.Writer.Header().Set("Content-Type", "application/json")
	//c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	a := c.PostForm("patient")
	bks := models.ModelsGetPatientOfArena(a)
	c.JSON(200, bks)


}


func MainJwt(c *gin.Context)  {
	//user := c.PostForm("user")
	//token := user.(jwt.Token)
	//
	//claims := token.Claims.(jwt.MapClaims)
	//
	//log.Println("User Name: ", claims["name"], "User ID: ", claims["jti"])
	c.String(http.StatusOK, "JWT метод")
}
func CheckToken(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	//c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	//user := c.PostForm("user")
	//token := user.(jwt.Token)
	//
	//claims := token.Claims.(jwt.MapClaims)
	//
	//log.Println("User Name: ", claims["name"], "User ID: ", claims["jti"])
	c.String(http.StatusOK, "JWT метод")
}

func Auth(c *gin.Context)  {
	name := c.PostForm("name")
	pass := c.PostForm("pass")
	checkUser := models.ModelsAuth(name, pass) //Проверка, есть ли пользователь в БД
	if checkUser[0].Count == "1" {

		userInfo := models.ModelsGetUserInfo(name, pass)
		token := auth.Login(c, userInfo) //Генерация jwt ключа пользователю
		c.JSON(http.StatusOK, gin.H{
			"message": "You were logged in!",
			"token": token,
		})
	} else {
		c.String(http.StatusForbidden, "Неверный логин, или пароль")
	}
}

