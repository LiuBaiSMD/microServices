package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"service-auth/handler"
	"service-auth/model"
	s "github.com/LiuBaiSMD/microServices/proto/auth-proto"
)

var (
	dockerMode string
)

func main() {

	// 初始化配置、数据库等信息
	//config.Init()

	// 使用consul注册
	//micReg := consul.NewRegistry(registryOptions)

	// 新建服务
	service := micro.NewService(
		micro.Name("tuyoo.micro.srv.auth"),
		//micro.Registry(micReg),
		micro.Version("latest"),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 初始化handler
			model.Init()
			// 初始化handler
			handler.Init()
		}),
	)

	// 注册服务
	s.RegisterServiceHandler(service.Server(), new(handler.Service))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
