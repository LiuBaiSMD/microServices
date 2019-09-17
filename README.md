## microServices说明
```markdown
本项目结合了多个go micro微服务，能够进行配置上传、用户信息管理（redis mysql数据库管理）、
```

```markdown
文件目录
.
├── README.md               项目说明
├── config                  各类型配置存储模块
├── dao                     数据库操作模块
├── go.mod                  
├── go.sum                  
├── handler                 逻辑处理模块
├── html                    html管理模块
├── service-auth            用户登录token验证服务
├── service-config-push     配置上传以及读取服务
├── service-web             web 请求处理服务 
└── util                    工具集
```

### ConfigService consul配置上传服务
```
1.配置上传
cd ConfigService
go run main.go loader.go  

2.配置读取
go client.go
```
### WebService 请求处理服务

```
cd WebService
go run main.go plugin.go

```

