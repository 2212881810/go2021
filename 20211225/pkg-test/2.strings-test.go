package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	v := runtime.Version()
	fmt.Println(v)
	first := strings.IndexByte(v, '.')
	last := strings.LastIndexByte(v, '.')
	fmt.Println(first)
	fmt.Println(last)

	r, err := strconv.ParseUint(v[first+1:last], 10, 64)

	if err != nil {
		fmt.Println(err)

	}

	fmt.Println(r)
}
