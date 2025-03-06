# 图片和视频下载代理服务

这是一个使用Go语言开发的Web服务，用于解决前端跨域下载图片和视频的问题。本服务作为代理，可以获取远程媒体资源并转发给客户端，同时处理CORS（跨域资源共享）问题。

## 特性

- 支持代理下载图片和视频文件
- 自动处理CORS跨域问题
- 简洁的Web界面用于测试和演示
- 纯Go实现，无外部依赖
- 支持自定义HTTP头
- 自动维护内容类型（Content-Type）

## 安装和运行

### 前提条件

- Go 1.16 或更高版本

### 安装Go（如果尚未安装）

#### macOS（使用Homebrew）
```bash
brew install go
```

#### Windows
从[Go官网](https://golang.org/dl/)下载安装包并安装。

#### Linux
```bash
# Ubuntu/Debian
sudo apt install golang

# Fedora
sudo dnf install golang
```

### 运行服务

1. 克隆或下载代码

2. 进入项目目录
```bash
cd server
```

3. 安装依赖
```bash
go mod tidy
```

4. 运行服务
```bash
go run cmd/proxy/main.go
```

服务默认在 http://localhost:8080 运行。

## 使用方法

### Web界面

访问 http://localhost:8080 使用Web界面下载媒体文件。

### API使用

#### 下载媒体文件
```
GET /api/download?url=媒体文件URL
```

#### 健康检查
```
GET /api/health
```

### 示例

#### 使用curl下载图片
```bash
curl -o image.jpg "http://localhost:8080/api/download?url=https://example.com/image.jpg"
```

#### 使用JavaScript在前端调用
```javascript
// 获取图片并显示
const img = document.createElement('img');
img.src = 'http://localhost:8080/api/download?url=' + encodeURIComponent('https://example.com/image.jpg');
document.body.appendChild(img);

// 下载视频
const a = document.createElement('a');
a.href = 'http://localhost:8080/api/download?url=' + encodeURIComponent('https://example.com/video.mp4');
a.download = 'video.mp4';
a.click();
```

## 环境变量

- `PORT`: 服务器监听端口，默认为8080

## 项目结构

```
server/
├── cmd/
│   └── proxy/
│       └── main.go         # 主程序入口
├── internal/
│   └── handlers/
│       ├── handlers.go     # 请求处理器
│       └── middleware.go   # 中间件
├── pkg/
│   └── downloader/
│       └── downloader.go   # 下载功能实现
├── static/                 # 静态文件目录
│   └── index.html         # Web界面
├── go.mod                  # Go模块定义
└── README.md              # 说明文档
```

## 许可证

MIT 