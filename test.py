import socket
from datetime import datetime

# 创建一个 TCP 连接
# HOST = 'qlmsmart.com'
HOST = "120.77.79.24"
PORT = 38082
sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
sock.connect((HOST, PORT))

# 向 TCP 服务器发送消息
message = 'Hello, TCP server!'
sock.sendall(message.encode())
while 1:
    # 接收 TCP 服务器的响应
    data = sock.recv(1024)
    timestamp = datetime.now().strftime('%Y-%m-%d %H:%M:%S')
    print(f"[{timestamp}] Received: {data.decode()}")

# 关闭 TCP 连接
sock.close()
