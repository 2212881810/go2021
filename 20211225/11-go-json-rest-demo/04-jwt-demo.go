package main

import (
	"log"
	"net/http"
	"time"

	// 注意哈， go-json-rest-middleware-jwt 没有更新，目前只能支持jwt-go v2版本
	"github.com/StephanDollberg/go-json-rest-middleware-jwt"
	"github.com/ant0ine/go-json-rest/rest"
)

func handle_auth(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(map[string]string{"authed": r.Env["REMOTE_USER"].(string)})
}

func main4() {
	jwt_middleware := &jwt.JWTMiddleware{
		Key:        []byte("secret key"),
		Realm:      "jwt auth",
		Timeout:    time.Hour,
		MaxRefresh: time.Hour * 24,
		Authenticator: func(userId string, password string) bool {
			return userId == "admin" && password == "admin"
		}}

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	api.Use(&rest.IfMiddleware{
		// Runtime condition that decides of the execution of IfTrue of IfFalse.
		Condition: func(r *rest.Request) bool {
			log.Printf("url path :%s\n", r.URL.Path)

			return r.URL.Path != "/login"
		},
		// 如果condition返回true，就会走jwt_middleware 这个中间件，然后就会调用LoginHandler，在LoginHandler中进行jwt鉴权
		IfTrue: jwt_middleware,
	})

	api_router, _ := rest.MakeRouter(
		// 如是
		rest.Post("/login", jwt_middleware.LoginHandler),

		rest.Get("/auth_test", handle_auth),
		rest.Get("/refresh_token", jwt_middleware.RefreshHandler),
	)
	api.SetApp(api_router)

	http.Handle("/api/", http.StripPrefix("/api", api.MakeHandler()))

	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		log.Fatal()
	}
}
