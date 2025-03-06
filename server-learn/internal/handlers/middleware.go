package handlers

import (
	"net/http"
)

// 跨域中间件
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 允许的请求来源
		w.Header().Set("Access-Control-Allow-Origin", "localhost:3000,insta-j.vercel.app")
		// 允许的请求方法
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		// 允许的请求头
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		// 暴露的请求头
		w.Header().Set("Access-Control-Expose-Headers", "Content-Disposition, Content-Length")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// 日志中间件
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("[INFO]", r.Method, r.URL.Path, "from", r.RemoteAddr)

		next.ServeHTTP(w, r)
	})
}
