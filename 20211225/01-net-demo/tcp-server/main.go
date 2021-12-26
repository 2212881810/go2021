package main

import (
	"fmt"
	"net"
)

func handConnection(conn *net.TCPConn) {
	// 这样就会循环使用这个conn,长连接
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)

		if err != nil {
			fmt.Println(err)
			break

		}

		fmt.Println(conn.RemoteAddr().String() + ":" + string(buf[:n]))
	}

}
func handConnection2(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read data from conn error :", err)
			return
		}

		fmt.Println(string(buf[:n]))
	}
}

func main() {

	li, _ := net.Listen("tcp", ":8080")

	for {
		conn, err := li.Accept()

		if err != nil {
			fmt.Println("acc error :", err)
			continue
		}
		go handConnection2(conn)
	}

	//tcpAddr, _ := net.ResolveTCPAddr("tcp", ":8080")
	//
	//ln, err := net.ListenTCP("tcp", tcpAddr)
	//
	//if err != nil {
	//	fmt.Println("listen error :", err)
	//	return
	//}
	//
	//for {
	//
	//	conn, err := ln.AcceptTCP()
	//
	//	if err != nil {
	//		fmt.Println("accept error :", err)
	//		continue
	//	}
	//	go handConnection(conn)
	//}
}
