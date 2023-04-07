package api

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
	"web/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DDNS struct {
	Domain string    `bson:"domain"`
	Date   string    `bson:"date"`
}

func GetApi() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	collection := database.GetDB().Database("test").Collection("DDNS")
	router.GET("/", func(c *gin.Context) {
		ip := strings.Split(c.Request.Header.Get("X-Forwarded-For"), ",")[0]
		if ip == "" {
			ip = c.Request.RemoteAddr
		}
		ip = strings.TrimSpace(ip)
	
		// 如果IP地址是IPv6地址，则使用SplitHostPort函数来分割IP地址和端口号
		if strings.Contains(ip, ":") {
			ip, _, _ = net.SplitHostPort(ip)
		}
	
		filter := bson.M{"domain": ip}
		var result DDNS
		err := collection.FindOne(context.Background(), filter).Decode(&result)
		if err != nil && err != mongo.ErrNoDocuments {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	
		if err == mongo.ErrNoDocuments {
			currentTime := time.Now().UTC()
			record := DDNS{Domain: ip, Date: currentTime.Format("2006-01-02 15:04:05")}
			_, err := collection.InsertOne(context.Background(), record)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
		}
	
		c.String(200,ip)
	})
	



	router.POST("/login", func(c *gin.Context) {
		param := make(map[string]interface{})
		err := c.BindJSON(&param)
			fmt.Println(param)
			fmt.Println(param["pwd"])
			if err != nil {
				return 
			}

	
	})

	router.Run("[::]:8081")
}


