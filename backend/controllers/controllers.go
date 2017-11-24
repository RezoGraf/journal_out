package controllers

import (
	"github.com/gin-gonic/gin"
	"../models"
	"../auth"
	"net/http"
	"fmt"
)

func FindPatientInArena(c *gin.Context)  {
	user := auth.EncodeJwt(c)
	fmt.Println(user["id"], user["name"])
	a := c.PostForm("patient")
	bks := models.ModelsGetPatientOfArena(a)
	c.JSON(200, bks)
}

func GetCurentUser(c *gin.Context)  { //Получение id и name по JWT token
	user := auth.EncodeJwt(c)
	c.JSON(200, gin.H{
		"id": user["id"],
		"name": user["name"],
	})
}


func MainJwt(c *gin.Context)  {
	c.String(http.StatusOK, "JWT метод")
}
func CheckToken(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
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

func AddPrivivka(c *gin.Context)  {
	encodeuser := auth.EncodeJwt(c)
	userId := 	 encodeuser["id"].(string)
	vaccination  := c.PostForm("vaccination")
	date		 := c.PostForm("date")
	preparat	 := c.PostForm("preparat")
	seria 		 := c.PostForm("seria")
	doza 		 := c.PostForm("doza")
	numberkart 	 := c.PostForm("numberkart")

	models.ModelsAddPrivivka(userId, vaccination, date, preparat, seria, doza, numberkart)
	c.JSON(200, "ok")
}

func GetPrivivka(c *gin.Context)  {
	type_vaccination := c.PostForm("type_vaccination")
	numberKart := c.PostForm("numberKart")
	fmt.Println(type_vaccination, "!!!! ", numberKart)
	bks := models.ModelsGetPrivivka(type_vaccination, numberKart)
	c.JSON(200, bks)
}

