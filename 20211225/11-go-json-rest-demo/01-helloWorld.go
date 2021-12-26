package main

import (
	"fmt"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

func main1() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	app := rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) {

		addr := r.RemoteAddr
		fmt.Printf("remote addr :%s \n", addr)
		//w.WriteHeader(200) // 好像就是响应码

		w.Header().Set("hello", "World!")
		w.WriteJson("{\"name\":\"admin\"}")
	})

	api.SetApp(app)
	http.ListenAndServe(":8088", api.MakeHandler()) // 调用了app
}
