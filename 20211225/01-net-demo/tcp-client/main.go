package main

import (
	"bufio"
	"net"
	"os"
)

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", ":8080")
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	// 标准输入，由控制台输入
	reader := bufio.NewReader(os.Stdin)
	for { // 这里是个死循环，所以前面创建的conn不会关闭，相当于是个长连接了
		line, _, _ := reader.ReadLine()
		conn.Write(line)
	}
}
