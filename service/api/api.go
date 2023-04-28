package api

import (

	"github.com/gin-gonic/gin"

)



func GetApi() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	
	router.GET("/",setip)

	router.POST("/login", login)

	router.GET("/login", page)

	router.Run("[::]:8081")
}
