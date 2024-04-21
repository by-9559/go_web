package tcp

import (
	"fmt"
	"net"
	"sync"
)

var connections []net.Conn
var mutex sync.Mutex

func TCP_go() {
	// 监听本地的 TCP 1234 端口
	ln, err := net.Listen("tcp", ":8082")
	if err != nil {
		panic(err)
	}

	defer ln.Close()

	fmt.Println("TCP server is running on port 8082")

	// 定义一个连接对象的切片和一个互斥锁

	for {
		// 等待客户端连接
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		// 将连接对象添加到切片中
		mutex.Lock()
		connections = append(connections, conn)
		mutex.Unlock()
		// 启动一个 goroutine 处理连接
		go handleConnection(conn, &connections, &mutex)
	}
}

func handleConnection(conn net.Conn, connections *[]net.Conn, mutex *sync.Mutex) {
	defer conn.Close()

	// 读取客户端发送的数据
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Connection closed:", conn.RemoteAddr())
			break
		}

		// 将收到的数据转成字符串并输出
		data := string(buffer[:n])
		// 向客户端发送响应数据
		response := "Hello, client!     " + conn.RemoteAddr().String()
		_, err = conn.Write([]byte(response))
		if err != nil {
			panic(err)
		}

		// fmt.Println("Response sent to client:", response)

		// 遍历连接对象的切片，向除了当前连接外的所有连接发送消息
		mutex.Lock()
		for _, c := range *connections {
			if c != conn {
				_, err = c.Write([]byte(data))
				if err != nil {
					fmt.Println("Failed to send data to:", c.RemoteAddr())
				}
			}
		}
		mutex.Unlock()
	}

	// 从切片中移除连接对象
	mutex.Lock()
	for i, c := range *connections {
		if c == conn {
			*connections = append((*connections)[:i], (*connections)[i+1:]...)
			break
		}
	}
	mutex.Unlock()
}

func Get_connections() *[]net.Conn {
	return &connections
}
