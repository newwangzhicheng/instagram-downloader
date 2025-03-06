package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"instagram-downloader/server/internal/handlers"
)

func main() {
	// 设置日志格式
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// 获取端口配置，默认为8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 创建路由复用器
	mux := http.NewServeMux()

	// 注册API路由
	mux.HandleFunc("/api/download", handlers.DownloadHandler)
	mux.HandleFunc("/api/health", handlers.HealthCheckHandler)

	// 添加静态文件服务
	// 获取当前工作目录
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("无法获取当前工作目录: %v", err)
	}

	// 静态文件目录路径
	staticDir := filepath.Join(cwd, "static")

	// 确保静态文件目录存在
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		os.MkdirAll(staticDir, 0755)
		// 创建一个简单的index.html文件
		indexHTML := `
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>图片和视频下载代理</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            max-width: 800px;
            margin: 0 auto;
            padding: 20px;
            line-height: 1.6;
        }
        .container {
            background-color: #f5f5f5;
            border-radius: 8px;
            padding: 20px;
            box-shadow: 0 2px 4px rgba(0,0,0,0.1);
        }
        input[type="url"] {
            width: 100%;
            padding: 8px;
            margin: 10px 0;
            border: 1px solid #ddd;
            border-radius: 4px;
        }
        button {
            background-color: #4CAF50;
            color: white;
            border: none;
            padding: 10px 15px;
            border-radius: 4px;
            cursor: pointer;
        }
        button:hover {
            background-color: #45a049;
        }
        pre {
            background-color: #f8f8f8;
            padding: 10px;
            border-radius: 4px;
            overflow-x: auto;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>图片和视频下载代理</h1>
        <p>输入图片或视频的URL，点击下载按钮解决跨域问题。</p>
        
        <input type="url" id="mediaUrl" placeholder="请输入图片或视频URL" />
        <button onclick="downloadMedia()">下载</button>
        
        <h2>API用法：</h2>
        <pre>GET /api/download?url=图片或视频的URL</pre>
        
        <h2>示例：</h2>
        <pre>https://example.com/api/download?url=https://example.com/image.jpg</pre>
    </div>

    <script>
        function downloadMedia() {
            const mediaUrl = document.getElementById('mediaUrl').value.trim();
            if (!mediaUrl) {
                alert('请输入有效的URL');
                return;
            }
            
            // 构建下载URL
            const downloadUrl = '/api/download?url=' + encodeURIComponent(mediaUrl);
            
            // 创建一个隐形的a标签并点击它来下载
            const a = document.createElement('a');
            a.href = downloadUrl;
            a.target = '_blank';
            document.body.appendChild(a);
            a.click();
            document.body.removeChild(a);
        }
    </script>
</body>
</html>
`
		err = os.WriteFile(filepath.Join(staticDir, "index.html"), []byte(indexHTML), 0644)
		if err != nil {
			log.Fatalf("无法创建index.html文件: %v", err)
		}
	}

	// 服务静态文件
	fs := http.FileServer(http.Dir(staticDir))
	mux.Handle("/", http.StripPrefix("/", fs))

	// 应用中间件
	handler := handlers.LoggerMiddleware(handlers.CORSMiddleware(mux))

	// 创建自定义的HTTP服务器
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	// 启动服务器
	fmt.Printf("服务器启动在 http://localhost:%s\n", port)
	log.Fatal(server.ListenAndServe())
}
