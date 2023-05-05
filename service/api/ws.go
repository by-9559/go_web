package api

import (
	"fmt"
	"log"
	"net/http"
	tcp "web/service/tcp"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var messages []string;

var wsCons []*websocket.Conn

func websocketHandler(c *gin.Context) {
	// 升级HTTP连接为WebSocket连接
	wsConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	wsCons = append(wsCons, wsConn)
	wsConn.WriteJSON(messages)
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

		connections := tcp.Get_connections()
		for _, c := range *connections {

			_, err = c.Write([]byte(p))
			if err != nil {
				fmt.Println("Failed to send data to:", c.RemoteAddr())
			}

		}

		for _, c := range wsCons {
			if c != wsConn {
				err = c.WriteJSON(messages)
				if err != nil {
					removeConnection(c)
					fmt.Println("Failed to send data to:", c.RemoteAddr())
				}
			}

		}
		// 打印接收到的消息
		log.Printf("Received message: %s\n", p)

		err = wsConn.WriteJSON(messages)
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
