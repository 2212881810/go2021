package main

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

func main3() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	api.Use(&rest.AuthBasicMiddleware{
		Realm: "127.0.0.1", //  域，请求域， 响应时会带此属性
		Authenticator: func(userId string, password string) bool {
			if userId == "admin" && password == "admin" {
				return true
			}
			return false
		},

		// 此方法必须要在 Authenticator 返回true之后才会进入！！！
		//Authorizator: func(userId string, request *rest.Request) bool {
		//	fmt.Printf("Authorizator: userId->%s\n", userId)
		//	return true
		//},
	})
	api.SetApp(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) {
		w.WriteJson(map[string]string{"Body": "Hello World!"})
	}))
	http.ListenAndServe(":8088", api.MakeHandler())

}
