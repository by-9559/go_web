package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

var addr = ""

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	target_url := fmt.Sprintf("https://%s:35212/", get_external())
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": target_url,
		})
	})

	router.POST("/ip", func(c *gin.Context) {
		addr = c.RemoteIP()
	})

	router.Run(":38080")
}

func get_external() string {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	return string(content)
}
