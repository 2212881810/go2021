package main

import (
	"fmt"
	"net/rpc"
	"time"
)

// 这两个结构体必须与server端一致
type Req struct {
	NumOne int
	NumTwo int
}

type Resp struct {
	Num int
}

func main() {
	req := Req{NumOne: 1, NumTwo: 2}
	var resp Resp
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return

	}

	// Call是同步方法,它的底层实现使用了Go方法+chan
	//err = client.Call("Server.Add", req, &resp)
	//if err != nil {
	//	fmt.Println("call server add error:", err)
	//	return
	//}
	//fmt.Println(resp)

	// done为nil,Go方法内部会为我们创建一个chan
	call := client.Go("Server.Add", req, &resp, nil)

	//fmt.Println("异常请求,我可以在这里做一些事情,等待server响应!")
	//<-call.Done
	//fmt.Println(resp)

	// 这样也可以哈！！！！！！！！！！！
	for {
		select {
		case <-call.Done:
			// 获取结果值 ，结事
			fmt.Println(resp)
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("我等着呢")

		}
	}
}
