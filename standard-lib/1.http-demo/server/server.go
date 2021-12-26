package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		addr := r.RemoteAddr
		fmt.Println(addr)
		w.Write([]byte("recv request~!"))

	})

	//err := http.ListenAndServe(":8080", http.DefaultServeMux)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
