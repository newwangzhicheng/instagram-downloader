package downloader

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"path"
	"strings"
	"time"
)

// ProxyDownload 代理下载指定URL的文件并转发到响应
func ProxyDownload(w http.ResponseWriter, fileURL string) error {
	// 创建带超时的HTTP客户端
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	
	// 创建请求
	req, err := http.NewRequest(http.MethodGet, fileURL, nil)
	if err != nil {
		return fmt.Errorf("创建请求失败: %w", err)
	}
	
	// 设置常见的请求头，模拟浏览器行为
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("发送请求失败: %w", err)
	}
	defer resp.Body.Close()
	
	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("远程服务器返回错误: %s", resp.Status)
	}
	
	// 从原始响应中获取内容类型
	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		// 如果内容类型为空，尝试从URL路径猜测
		ext := path.Ext(fileURL)
		if ext != "" {
			contentType = mime.TypeByExtension(ext)
		}
		
		// 如果仍然为空，使用通用二进制类型
		if contentType == "" {
			contentType = "application/octet-stream"
		}
	}
	
	// 获取文件名
	filename := getFilenameFromURL(fileURL)
	
	// 设置响应头
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	w.Header().Set("Access-Control-Allow-Origin", "*") // 允许跨域访问
	
	// 如果原始响应有Content-Length，也设置它
	if resp.Header.Get("Content-Length") != "" {
		w.Header().Set("Content-Length", resp.Header.Get("Content-Length"))
	}
	
	// 将文件内容复制到响应
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		return fmt.Errorf("传输数据失败: %w", err)
	}
	
	return nil
}

// getFilenameFromURL 从URL中提取文件名
func getFilenameFromURL(fileURL string) string {
	// 提取路径的最后一部分作为文件名
	urlPath := path.Base(fileURL)
	
	// 移除查询参数
	if idx := strings.IndexByte(urlPath, '?'); idx >= 0 {
		urlPath = urlPath[:idx]
	}
	
	// 确保文件名不为空
	if urlPath == "" || urlPath == "." {
		return "download"
	}
	
	return urlPath
} 