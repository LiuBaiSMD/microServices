# Own Station 服务

这是一个处理用户登录请求等业务的项目，主要功能如下
```
1.用户注册
2.用户登录
3.用户修改密码
4.配置读取
5.用户token管理（简单版本）
6.用户使用userToken即可进行登录
```

## 运行环境
```
0.go micro
1.redis服务
2.mysql服务
3.consul服务
```
## 启动方式
```cassandraql
1.启动配置程序，在consul中添加配置，添加配置方法可参考配置上传项目
```
[配置上传项目](https://github.com/LiuBaiSMD/start-Go-micro/tree/master/dockerPRJ/dockerComp/consul-manager)

```
2.命令行启动mysql服务，创建一个micro_user数据库

mysql.server start
CREATE DATABASE micro_user; mysql服务中执行
```
```
3.命令行启动redis服务，本地端口默认配置即可，localhost:6379	

redis-server &  #后台运行redis服务
```
```
4.运行服务
go run main.go plugins.go
```

## 项目注意
```
即将改进的地方
1.自动创建micro_user数据库
2.consul配置连接不上时，能够自己创建consul服务，上传配置，以供本项目正常运行
3.增加用户登录后的服务
4.密码以及密码检查更加规范
5.使用更加灵活的初始化方法，或许从consul中读取redis、mysql的连接配置
```

```
特点：
1.账户每次登录会重新使用密码登录会重新生成token
```