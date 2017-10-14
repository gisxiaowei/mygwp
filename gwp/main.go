package main

import (
	"net/http"
	"time"
)

func main() {
	p("GWP", version(), "端口", config.Address)

	// 处理静态资源
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// 启动服务器
	server := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
