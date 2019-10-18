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
	"service-web/handler"
	"github.com/LiuBaiSMD/microServices/base/myLog"
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
		web.Name("websocket"),
		web.Version("latest"),
		web.Registry(micReg),
		web.Address(":8080"),
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
	// 注册登录接口
	handler.SetHandle(service, handler.WebHandlerConfig)
	handler.SetHandleFunc(service, handler.WebConfig)
	myLog.SetReport(handler.Report)
	myLog.Logger.Info("test123")
	myLog.Logger.Error("我来看看有没有错误!")
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("over")
}

func registryOptions(ops *registry.Options) {
	dockerMode := os.Getenv("RUN_DOCKER_MODE")
	if dockerMode == "on"{
		log.Log("docker模式")
		ops.Addrs = []string{"consul2"}
	}else{
		log.Log("本地模式")
		ops.Addrs = []string{"127.0.0.1:8500"}
	}
}





