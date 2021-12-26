package main

import (
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	"time"
)
// 以流的方式一直向客户端发送数据
func main5() {

	api := rest.NewApi()
	api.Use(&rest.AccessLogApacheMiddleware{})

	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(rest.Get("/stream", StreamThings))

	if err != nil {
		log.Printf("create router error : %v\n", err)
		return
	}

	api.SetApp(router)
	http.ListenAndServe(":8088", api.MakeHandler())

}

type Thing struct {
	Name string
}

func StreamThings(w rest.ResponseWriter, r *rest.Request) {
	cpt := 0
	for {
		cpt++
		w.WriteJson(&Thing{
			Name: fmt.Sprintf("thing #%d", cpt),
		})
		// 为什么w能强转成http.ResponseWriter， 因为w实现了 http.ResponseWriter 的三个方法
		w.(http.ResponseWriter).Write([]byte("\n"))

		// flush the buffer data to client
		// 为什么w能强转成http.Flusher， 因为w实现了 http.Flusher 的flush方法
		w.(http.Flusher).Flush()

		time.Sleep(time.Duration(3) * time.Second)
	}
}
