package main

import (
	"context"
	"fmt"
	"time"
)

//使用context来做到父上下文控制子上下文的目的
func main() {

	// 可以将flag换成context来实现
	//flag := make(chan bool)

	// context的第1种使用方式(第一个api WithCancel)
	// WithCancel方法内部创建一个Done channel,当调用clear方法时,Done channel会关闭
	//ctx, clear := context.WithCancel(context.Background())

	// context的第2种使用方法,携带值
	ctx := context.WithValue(context.Background(), "time", "时间到了")
	ctx, clear := context.WithCancel(ctx)

	// 当前时间+10秒后到期
	//ctx, clear = context.WithDeadline(context.Background(),time.Now().Add(time.Second*10))
	// 10秒后到期
	//ctx, clear = context.WithTimeout(context.Background(), time.Second*10)

	message := make(chan int)

	go son(ctx, message)

	for i := 0; i < 10; i++ {
		message <- i
	}

	//flag <- true
	clear()
	time.Sleep(time.Second)
	fmt.Println("主进程结束！")
}

func son(ctx context.Context, message chan int) {
	// 创建一个定时器， 一秒钟执行一次
	tick := time.Tick(time.Second)
	for _ = range tick {
		select {
		case m := <-message:
			fmt.Printf("%d\n", m)
		case <-ctx.Done(): // 主方法中调用clear方法时,这里才有反应
			//fmt.Printf("我结束了！\n")
			fmt.Println("我结束了:", ctx.Value("time"))
			return
		}

	}
}

//func son(flag chan bool, message chan int) {
//	// 创建一个定时器， 一秒钟执行一次
//	tick := time.Tick(time.Second)
//	for _ = range tick {
//		select {
//		case m := <-message:
//			fmt.Printf("%d\n", m)
//		case <-flag:
//			fmt.Printf("我结束了！\n")
//			return
//		}
//
//	}
//}
