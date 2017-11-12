package main

import (
	"./controllers"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/jwt"
	"./auth"
)

var (
	mysupersecretpassword = "mySecret"
)

func main() {

	router := gin.Default()

	jwtGroup := router.Group("/jwt")
	jwtGroup.Use(jwt.Auth(mysupersecretpassword))

	jwtGroup.GET("/main", controllers.MainJwt)
	jwtGroup.POST("/FindPatientInArena", controllers.FindPatientInArena)
	jwtGroup.POST("checktoken", controllers.CheckToken)
	//router.Static("/", "./public")
	//router.POST("/upload", controllers.Xmlparser)
	//router.POST("/get_data", controllers.GetXlsxFile)
	//router.GET( "/get_data", controllers.DownloadGetXlsxFile)

	router.GET("/login", auth.Login)

	router.Run(":8084")
}
