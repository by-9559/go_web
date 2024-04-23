package api

import (
	"fmt"
	"log"
	"net/http"
	"sync"

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
var channel = "chat"
var wsCons []*websocket.Conn
var mutex sync.Mutex

type WebSocketConn struct {
	ID string
	// Request *http.Request
	Conn *websocket.Conn
}

func getChatList() []string {
	redis.LTrim(channel, -300, -1).Result()
	messages, _ := redis.LRange(channel, 0, -1).Result()
	return messages
}

func websocketHandler(c *gin.Context) {
	wsConn, _ := upgrader.Upgrade(c.Writer, c.Request, nil)
	wsConn.WriteJSON(getChatList())
	wsCons = append(wsCons, wsConn)
	defer wsConn.Close()

	for {
		// 读取客户端发送的消息
		_, p, err := wsConn.ReadMessage()
		if err != nil {

			return
		}
		mutex.Lock()
		go redis.RPush(channel, p).Err()
		
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

		log.Printf("Received message: %s\n", p)

		err = wsConn.WriteJSON(getChatList())
		if err != nil {
			// 处理错误
			return
		}
		mutex.Unlock()

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
