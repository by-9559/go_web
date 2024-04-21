package api

import (
	"github.com/gin-gonic/gin"
)

func GetApi() {

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", setip)

	r.POST("/login", login)

	r.GET("/login", page)

	r.GET("/favicon.ico", func(c *gin.Context) { c.File("./templates/favicon.ico") })

	r.GET("/send", tcp_send)

	r.GET("/ws", websocketHandler)

	r.GET("/getTCPConns", tcplist)

	r.Run("[::]:8081")
}
