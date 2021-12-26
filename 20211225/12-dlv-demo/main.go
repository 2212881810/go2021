package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var a = 123
	fmt.Printf("%s\n", a)

	ch := make(chan struct{}, 2)
	go func() {
		for {
			time.Sleep(100 * time.Second)
			fmt.Println(os.Getgroups())
		}
	}()

	go func() {
		for {
			time.Sleep(10 * time.Second)
			fmt.Println(os.Getgroups())
		}
	}()

	<-ch
	<-ch

}
