package api

import (
	"fmt"
	"log"
	"net/http"
	db "web/database"
	tcp "web/service/tcp"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var redis = db.GetRedis()

var wsCons []*websocket.Conn

func getChatList() []string {
	redis.LTrim("chat", -300, -1).Result()
	messages, _ := redis.LRange("chat", 0, -1).Result()
	return messages
}

func websocketHandler(c *gin.Context) {
	// 升级HTTP连接为WebSocket连接
	wsConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	wsCons = append(wsCons, wsConn)

	wsConn.WriteJSON(getChatList())
	if err != nil {
		// 处理错误
		return
	}
	defer wsConn.Close()

	for {
		// 读取客户端发送的消息
		_, p, err := wsConn.ReadMessage()
		if err != nil {
			// 处理错误
			return
		}

		err = redis.RPush("chat", p).Err()
		if err != nil {
			panic(err)
		}
		connections := tcp.Get_connections()
		for _, c := range *connections {

			_, err = c.Write([]byte(p))
			if err != nil {
				fmt.Println("Failed to send data to:", c.RemoteAddr())
			}

		}

		for _, c := range wsCons {
			if c != wsConn {
				err = c.WriteJSON(getChatList())
				if err != nil {
					removeConnection(c)
					fmt.Println("Failed to send data to:", c.RemoteAddr())
				}
			}

		}
		// 打印接收到的消息
		log.Printf("Received message: %s\n", p)

		err = wsConn.WriteJSON(getChatList())
		if err != nil {
			// 处理错误
			return
		}

	}
}

func removeConnection(conn *websocket.Conn) {
	for i, c := range wsCons {
		if c == conn {
			wsCons = append(wsCons[:i], wsCons[i+1:]...)
			break
		}
	}
	// 处理连接关闭...
}
