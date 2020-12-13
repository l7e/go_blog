package main

import (
	"fmt"
	setting "go_blog/pkg"
	"go_blog/routers"
	"net/http"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.ReadTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
