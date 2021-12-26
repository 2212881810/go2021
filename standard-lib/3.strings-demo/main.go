package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "dfdjhf:djfgdjgfd:dfhdk"
	n := strings.SplitN(s, ":", 2) //  2表示切片中只有两个元素，3：表示切片中有3个元素

	fmt.Printf("%v\n", n)

	r := strings.Split(s, ":") // 无第三个参数， 全切
	fmt.Printf("%v\n", r)
}
