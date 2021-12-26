package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Server struct {
}

type Req struct {
	NumOne int
	NumTwo int
}

type Resp struct {
	Num int
}

// 方法必须这个声明,
// Add是Server的一个方法
func (s *Server) Add(req Req, resp *Resp) error {
	time.Sleep(5 * time.Second)
	resp.Num = req.NumOne + req.NumTwo
	return nil
}

func main() {
	server := new(Server)
	// 注册服务
	rpc.Register(server)
	rpc.HandleHTTP()

	listen, e := net.Listen("tcp", ":8080")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(listen, nil)
}
