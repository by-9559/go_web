package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var messages []string

func websocketHandler(c *gin.Context) {
	// 升级HTTP连接为WebSocket连接
	wsConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
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

		// 将消息保存到切片中
		messages = append(messages, string(p))

		// 打印接收到的消息
		log.Printf("Received message: %s\n", p)

		err = wsConn.WriteJSON(messages)
		if err != nil {
			// 处理错误
			return
		}

	}
}
