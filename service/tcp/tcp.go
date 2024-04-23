package tcp

import (
	"fmt"
	"net"
	"sync"
)

var (
	connections []net.Conn
	mutex       sync.Mutex
)

// TCPGo 启动TCP服务器监听
func TCPGo(port int) {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d",port))
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	fmt.Printf("TCP server is running on port %d\n", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		mutex.Lock()
		connections = append(connections, conn)
		mutex.Unlock()

		go handleConnection(conn)
	}
}

// handleConnection 处理每个客户端连接
func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Connection closed by remote:", conn.RemoteAddr())
			break
		}

		data := string(buffer[:n])
		response := "Hello, client! " + conn.RemoteAddr().String()
		if _, err = conn.Write([]byte(response)); err != nil {
			fmt.Println("Error sending response to client:", conn.RemoteAddr())
		}

		sendMessageToOtherClients(conn, data)
	}

	removeConnection(conn)
}

// sendMessageToOtherClients 向其他客户端发送消息
func sendMessageToOtherClients(sender net.Conn, data string) {
    mutex.Lock()
    defer mutex.Unlock()

    var failedConns []net.Conn  // 存储发送失败的连接

    for _, conn := range connections {
        if conn != sender {
            if _, err := conn.Write([]byte(data)); err != nil {
                fmt.Println("Failed to send data to other client:", conn.RemoteAddr())
                failedConns = append(failedConns, conn)
            }
        }
    }

    // 处理发送失败的连接
    for _, failedConn := range failedConns {
        removeConnection(failedConn)
        failedConn.Close()  // 关闭失败的连接
    }
}

// removeConnection 从连接列表中移除指定的连接
func removeConnection(conn net.Conn) {
    for i, c := range connections {
        if c == conn {
            connections = append(connections[:i], connections[i+1:]...)
            return
        }
    }
}

// GetConnections 返回当前的连接列表
func GetConnections() *[]net.Conn {
	return &connections
}
