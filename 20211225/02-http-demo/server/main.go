package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler(res http.ResponseWriter, req *http.Request) {
	m := req.Method

	switch m {
	case "GET":
		res.Write([]byte("接收到请求GET"))
		break
	case "POST":
		// 获取请求体中的内容
		body := req.Body
		// 读
		b, _ := ioutil.ReadAll(body)
		fmt.Println("recv content:", string(b))
		// 设置响应头
		res.Header().Add("test", "testHeader")
		res.WriteHeader(http.StatusOK)
		// 设置响应内容
		res.Write(b)
	}
}
func main() {
	// 设置请求路径是http://127.0.0.2:8080/test
	http.HandleFunc("/test", handler)
	// handler为空时，默认调用 http.DefaultServeMux
	//http.ListenAndServe("127.0.0.2:8080", nil)

	//mux := http.NewServeMux()
	////效果跟上传传nil是一样的，但是我们要实现mux的方法，做自己想做的事
	http.ListenAndServe("127.0.0.2:8080", http.DefaultServeMux)

}
