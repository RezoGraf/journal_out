package controllers

import (
	"github.com/gin-gonic/gin"
	"../models"
	"../auth"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin/render"
	"context"
)

func FindPatientInArena(c *gin.Context)  {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	//c.Writer.Header().Set("Content-Type", "application/json")
	//c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	a := c.PostForm("patient")
	fmt.Println("controller ", a)
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
	fmt.Println("!!!!", name, pass)
	checkUser := models.ModelsAuth(name, pass)
	fmt.Println(checkUser[0].Count)
	if checkUser[0].Count == "1" {
		fmt.Println(checkUser[0].Count)
		token := auth.Login(c)
		c.JSON(http.StatusOK, gin.H{
			"message": "You were logged in!",
			"token": token,
		})
	} else {
		c.String(http.StatusForbidden, "Неверный логин, или пароль")
	}
}

