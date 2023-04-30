package api

import (
	"context"
	"fmt"

	"net"
	"net/http"
	"strings"
	"time"

	"web/service/tcp"
	db "web/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DDNS struct {
	Domain string `bson:"domain"`
	Date   string `bson:"date"`
}

var client = db.GetDB()
var collection = client.Database("test").Collection("DDNS")

func login(c *gin.Context) {
	param := make(map[string]interface{})
	err := c.BindJSON(&param)
	fmt.Println(param)
	fmt.Println(param["pwd"])
	if err != nil {
		return
	}
}

func page(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "极简生活",
	})
}

func setip(c *gin.Context) {
	ip := strings.Split(c.Request.Header.Get("X-Forwarded-For"), ",")[0]
	if ip == "" {
		ip = c.Request.RemoteAddr
		ip, _, _ = net.SplitHostPort(ip)
	}
	ip = strings.TrimSpace(ip)

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

	c.String(200, ip)
}

type ConnectionPoolStatus struct {
    MaxConnections int
    NumConnections int
    Connections    []string
}

func tcp_send(c *gin.Context) {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	str := ""
	for _, conn := range *tcp.Get_connections() {
		// 使用 conn 对象进行操作
		fmt.Printf("对象:%s\r\n", conn.RemoteAddr())

		_, err = conn.Write([]byte("你好天才"))
	
		str = conn.RemoteAddr().String()
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(str)
	c.String(http.StatusOK, "Message sent to TCP server: %s", str)
}
