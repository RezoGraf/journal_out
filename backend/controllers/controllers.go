package controllers

import (
	"github.com/gin-gonic/gin"
	"../models"
	"fmt"
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
