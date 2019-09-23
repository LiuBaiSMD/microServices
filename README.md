## microServices说明
```markdown
本项目结合了多个go micro微服务，能够进行配置上传、用户信息管理（redis mysql数据库管理）、
```

```markdown
.
├── README.md               项目说明
├── config                  各类型配置存储模块
├── go.mod
├── go.sum
├── proto                   各项目的proto库
├── service-auth            登录token的rpc服务（doing）
├── service-config-push     配置上传以及读取服务
├── service-discuss-room    web 讨论功能聊天室
├── service-web             web 用户服务
└── util                    工具集
```

### service-config-push consul配置上传服务
```markdown
1.配置上传
cd service-config-push
go run main.go loader.go  

2.配置读取测试
go client.go
```
### service-web 用户服务

```markdown
cd service-web
go run main.go plugin.go
```

### service-discuss-room 讨论功能聊天室服务
```markdown
cd service-discuss-room
go run main.go plugin.go  
```

## docker部署
```markdown
使用docker-compose.yml文件将服务一键部署到docker中
```
```markdown
运行步骤
1.下载安装docker
2.在每个service项目下使用go mod vendor 拉取依赖代码
3.docker-compose up  # 一键启动docker服务
4.docker ps 查看容器运行状态
```