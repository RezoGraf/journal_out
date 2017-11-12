package controllers

import (
	"github.com/gin-gonic/gin"
	"../models"
	"fmt"
	"net/http"
)

func FindPatientInArena(c *gin.Context)  {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	//c.Writer.Header().Set("Content-Type", "application/json")
	//c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	a := c.PostForm("patient")
	fmt.Println("controller ", a)
	bks := models.GetPatientOfArena(a)
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
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	//user := c.PostForm("user")
	//token := user.(jwt.Token)
	//
	//claims := token.Claims.(jwt.MapClaims)
	//
	//log.Println("User Name: ", claims["name"], "User ID: ", claims["jti"])
	c.String(http.StatusOK, "JWT метод")
}