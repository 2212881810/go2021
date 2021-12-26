package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(rest.Get("/message.txt", func(w rest.ResponseWriter, r *rest.Request) {

		w.Header().Set("Content-Type", "text/plain")
		w.(http.ResponseWriter).Write([]byte("Hello World."))
	}))

	if err != nil {
		log.Printf("create router error :%v\n",err)
		return
	}

	api.SetApp(router)
	http.ListenAndServe(":8088",api.MakeHandler())


}
