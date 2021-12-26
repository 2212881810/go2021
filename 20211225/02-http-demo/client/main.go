package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// 1. 创建客户端
	//client := new(http.Client)
	//// 2. 创建请求
	//request, _ := http.NewRequest("GET", "http://127.0.0.2:8080/test", nil)
	//// 3. 发送请求
	//response, _ := client.Do(request)
	//body := response.Body
	//b, _ := ioutil.ReadAll(body)
	//fmt.Println(string(b))

	c := http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}

	fmt.Println(c)
	// 1. 创建客户端
	client := new(http.Client)
	// 2. 创建请求
	var buf = []byte("{\"test\":\"post method\"}")

	request, _ := http.NewRequest("POST", "http://127.0.0.2:8080/test", bytes.NewBuffer(buf))
	// 3. 发送请求
	response, _ := client.Do(request)
	body := response.Body
	b, _ := ioutil.ReadAll(body)
	fmt.Println(string(b))

}
