# dockit

Dockit 是一个用Go语言编写的跨平台构建工具。它可以在不同的操作系统上编译出64位的可执行程序。

## 特性

特性

1. 支持Windows, Linux, macOS等主流操作系统
2. 无需安装任何依赖，直接编译即可
3. 编译过程简单快捷
4. 可定制化编译参数

## 安装
您可以直接下载编译好的可执行程序,也可以自行编译:
### 下载预编译版本
您可以在 [releases](https://github.com/youz88/dockit/releases) 页面下载适用于您操作系统的可执行程序。
### 自行编译
```shell
# 编译Windows 64位程序
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
# 编译Linux 64位程序 
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
# 编译macOS 64位程序
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
```

## 用法
```shell
Usage:
  dockit [command]

Available Commands:
  clear       Remove the docker image and container
  help        Help about any command
  launch      Pull the docker image and start the container
```
更多用法可以参考 dockit --help 查看。
