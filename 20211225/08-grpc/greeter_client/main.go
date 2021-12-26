package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "test_module/2021/08-grpc/pb"

	"google.golang.org/grpc"
)

const (
	defaultName = "郑钦锋"
)

var (
	// flag库搞出来的都是指针
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	//fmt.Println(*addr)
	//fmt.Println(*name)

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	greeterClient := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 必须传地址吗？？
	resp, err := greeterClient.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", resp.GetMessage())
}
