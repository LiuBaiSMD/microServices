
package main

import (
	"fmt"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
	"handlerManageTest/handler"
	//"github.com/gorilla/websocket"
)
func main() {

	// 初始化配置
	//base.Init()
	//filepath.Walk()
	// 使用consul注册
	//micReg := consul.NewRegistry(registryOptions)
	// 创建新服务
	handler.Init()
	service := web.NewService(
		web.Name("websocket-func"),
		web.Version("latest"),
		web.Address(":8081"),
	)

	// 初始化服务
	if err := service.Init(
	); err != nil {
		log.Fatal(err)
	}
	// 注册登录接口

	handler.HandlerFromConf(service, "manageFunc.json")

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("over")
}

