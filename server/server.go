package server

import (
	"TestTask/store"
	"fmt"
	"github.com/valyala/fasthttp"
)

var Default *fasthttp.Server

func NewServer(config *Config, store *store.Store) (server *fasthttp.Server, err error) {
	var router = NewRouter(store.Users)
	server = &fasthttp.Server{Handler: fasthttp.CompressHandler(router.Handler)}

	go func() {
		if err := server.ListenAndServe("localhost:8080"); err != nil {
			panic(err)
		}
	}()
	fmt.Print("All okey")
	return
}
