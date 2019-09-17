package main

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/log"
	"os"
	"time"
)

func main() {
	Init()
	// 注册一个consul服务
	micReg := consul.NewRegistry(registryOptions)
	// 新建服务
	service := micro.NewService(
		micro.Name("bambooRat.micro.srv.config"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// 服务初始化
	fmt.Println("config init")
	service.Init()

	// 启动服务
	fmt.Println("config run")
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	dockerMode := os.Getenv("RUN_DOCKER_MODE")
	if dockerMode != "on"{
		fmt.Println("本地模式1")
		ops.Timeout = time.Second * 5
		//ops.Addrs = []string{"consul4"}
		ops.Addrs = []string{consulConfigCenterAddr}
	}else{
		fmt.Println("docker模式")
		ops.Timeout = time.Second * 5
		ops.Addrs = []string{"consul4:8500"}
	}
}
