package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"instagram-downloader/server-learn/internal/handlers"
)

func main() {
	/** 设置日志格式
		LstdFlags： YYYY/MM/DD HH:MM:SS
		Lshortfile: 文件名和行号
		按位或，代表两个日志合并
	 **/
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// 获取端口，默认为8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 创建路由器
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handlers.HealthCheckHandler)

	// 添加中间件
	handler := handlers.CORSMiddleware(
		handlers.LoggerMiddleware(mux),
	)

	// 创建自定义的HTTP服务器
	// 这是一个结构体，所以取指针
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	// 启动HTTP服务器
	fmt.Printf("服务器启动 http://localhost:%s\n", port)
	log.Fatal(server.ListenAndServe())
}
