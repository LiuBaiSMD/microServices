<a title="Hits" target="_blank" href="https://github.com/xuyiwenak/consul-config-push"><img src="https://hits.b3log.org/b3log/hits.svg"></a>
# consul-config-push

单独的推送server推送配置到consul配置中心

## 目录结构  
```
├── LICENSE
├── README.md
├── client.go  // 模拟客户端，用户服务开发注册consul服务，读取配置的demo
├── conf
│   └── micro.yml // 静态配置
├── go.mod
├── go.sum
├── loader.go // 向配置中心推送配置
├── main.go
└── vendor
```  

## 使用解释
方便docker环境编译
```
go mod tidy
go mod vendor
```
### 1. 调试consul server
需要启动consul
```
consul agent -dev
// 或者从docker环境启动consul镜像
docker run <consul container name>
```
启动consul server
```
go run loader.go main.go
```
docker打包运行consul server
```
docker build -t consul-config-push .
docker run --rm -d consul-config-push
```
### 2. 调试consul client
启动consul cilent  
```
go run client.go
```
## 静态配置
conf 目录下的micro.yml 文件变更会自动检测到执行对应的逻辑




