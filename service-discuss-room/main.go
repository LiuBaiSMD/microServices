package main

import (
"fmt"
"github.com/gorilla/websocket"
"github.com/micro/cli"
"github.com/micro/go-micro/registry"
"github.com/micro/go-micro/registry/consul"
"github.com/micro/go-micro/util/log"
"github.com/micro/go-micro/web"
"net/http"
"os"
//"github.com/micro/go-web"
"microServices/handler"
)



var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {

	// 初始化配置
	//base.Init()

	// 使用consul注册
	micReg := consul.NewRegistry(registryOptions)
	// 创建新服务
	service := web.NewService(
		// 后面两个web，第一个是指是web类型的服务，第二个是服务自身的名字
		web.Name("discuss-server"),
		web.Version("latest"),
		web.Registry(micReg),
		web.Address(":8081"),
	)

	// 初始化服务
	if err := service.Init(
		web.Action(
			func(c *cli.Context) {
				// 初始化handler
				handler.Init()
			}),
	); err != nil {
		log.Fatal(err)
	}
	service.HandleFunc("/discussSub", handler.SetDiscuss)
	service.HandleFunc("/discussGet", handler.GetDiscuss)
	// 注册登录接口
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("over")
}

func registryOptions(ops *registry.Options) {
	dockerMode := os.Getenv("RUN_DOCKER_MODE")
	if dockerMode == "on"{
		log.Log("docker模式")
		ops.Addrs = []string{"consul1"}
	}else{
		log.Log("本地模式")
		ops.Addrs = []string{"127.0.0.1:8500"}
	}
}






