package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORS 中间件
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(http.StatusNoContent)
            return
        }

        c.Next()
    }
}


func GetApi(port int) {

	r := gin.Default()
	r.Use(CORSMiddleware())
	r.LoadHTMLGlob("templates/*")

	r.GET("/", setip)

	r.POST("/login", login)

	r.GET("/login", page)

	r.GET("/favicon.ico", func(c *gin.Context) { c.File("./templates/favicon.ico") })

	r.GET("/send", tcp_send)

	r.GET("/ws", websocketHandler)

	r.GET("/getTCPConns", tcplist)

	r.Run(fmt.Sprintf("[::]:%d",port))
}
