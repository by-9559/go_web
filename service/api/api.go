package api

import (
	"github.com/gin-gonic/gin"
	"log"
)

func GetApi() {

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.Use(gin.LoggerWithWriter(log.Writer()), gin.Recovery())

	r.GET("/", setip)

	r.POST("/login", login)

	r.GET("/login", page)

	r.GET("/favicon.ico", func(c *gin.Context) { c.File("./templates/favicon.ico") })

	r.GET("/send", tcp_send)

	r.Run("[::]:8081")
}
