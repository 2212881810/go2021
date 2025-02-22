package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main1() {
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		// gin.H 其实就是一个map[string]interface
		context.JSON(http.StatusOK, gin.H{"message": "pong"})

	})

	/*
		Run方法底层就是调用 http.ListenAndServe(address, engine)

	*/
	r.Run() // 源码里面默认就是8080端口

	//http.ListenAndServe(":8081", &MyHandler{
	//	Name: "郑钦锋",
	//})

	// this is my test git code

}

type MyHandler struct {
	Name string
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("%v\n", r)
	w.Write([]byte("test,,,,,,,,"))
}
