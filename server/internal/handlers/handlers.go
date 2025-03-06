package handlers

import (
	"encoding/json"
	"instagram-downloader/server/pkg/downloader"
	"net/http"
	"net/url"
)

// HealthCheckHandler 提供健康检查端点
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// 设置响应类型为JSON
	w.Header().Set("Content-Type", "application/json")
	
	// 返回简单的状态信息
	response := map[string]string{
		"status": "ok",
		"message": "服务正常运行",
	}
	
	// 编码为JSON并返回
	json.NewEncoder(w).Encode(response)
}

// DownloadHandler 处理媒体文件下载请求
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	// 只允许GET请求
	if r.Method != http.MethodGet {
		http.Error(w, "只支持GET请求", http.StatusMethodNotAllowed)
		return
	}
	
	// 获取URL参数
	urlParam := r.URL.Query().Get("url")
	if urlParam == "" {
		http.Error(w, "缺少url参数", http.StatusBadRequest)
		return
	}
	
	// 验证URL是否有效
	_, err := url.ParseRequestURI(urlParam)
	if err != nil {
		http.Error(w, "无效的URL", http.StatusBadRequest)
		return
	}
	
	// 下载文件
	err = downloader.ProxyDownload(w, urlParam)
	if err != nil {
		http.Error(w, "下载失败: "+err.Error(), http.StatusInternalServerError)
		return
	}
	
	// 注意: 实际的响应内容和头信息会在downloader.ProxyDownload函数中设置
} 