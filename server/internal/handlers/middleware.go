package handlers

import (
	"net/http"
)

// CORSMiddleware 是一个处理跨域请求的中间件
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 设置允许的请求来源
		w.Header().Set("Access-Control-Allow-Origin", "*")
		
		// 设置允许的HTTP方法
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		
		// 设置允许的请求头
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		// 设置暴露给客户端的响应头
		w.Header().Set("Access-Control-Expose-Headers", "Content-Disposition, Content-Length")
		
		// 处理预检请求
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		// 继续处理请求
		next.ServeHTTP(w, r)
	})
}

// LoggerMiddleware 是一个请求日志中间件
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 记录请求信息
		println("[INFO]", r.Method, r.URL.Path, "来自", r.RemoteAddr)
		
		// 继续处理请求
		next.ServeHTTP(w, r)
	})
} 