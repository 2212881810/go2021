package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

func main2() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(rest.Get("/lookup/#host", func(writer rest.ResponseWriter, request *rest.Request) {
		host := request.PathParam("host")
		fmt.Printf("req pathparam:%v\n", host)
		ip, err := net.LookupIP(host)
		if err != nil {
			//我草， 直接就写到响应流中去了
			rest.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.WriteJson(&ip)
	}))

	if err != nil {
		log.Fatal(err)
		return
	}
	api.SetApp(router)
	err = http.ListenAndServe(":8088", api.MakeHandler())
	if err != nil {
		fmt.Println(err)

	}

}
