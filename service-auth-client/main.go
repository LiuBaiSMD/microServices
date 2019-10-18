package main

import (
	"context"
	"fmt"
	model "github.com/LiuBaiSMD/microServices/proto/auth"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"os"
	"github.com/micro/go-micro/transport"
	//"github.com/micro/go-micro/util/log"
)

func main() {
	// 我这里用的etcd 做为服务发现
	var consulAddr string
	dockerMode := os.Getenv("RUN_DOCKER_MODE")
	if dockerMode == "on"{
		consulAddr = "consul2:8500"
	}else{
		consulAddr = "127.0.0.1:8500"
	}
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			consulAddr,
		}
	})
	tsp := transport.NewTransport(transport.Addrs("127.0.0.1:8500"))
	//l, err := tsp.Dial("127.0.0.1:39707")
	//if err != nil {
	//	fmt.Println("Unexpected listen err: %v", err)
	//}
	//defer l.Close()
	//初始化服务
	service := micro.NewService(
		micro.Name("auth-client"),
		micro.Registry(reg),
		micro.Version("latest"),
		micro.Transport(tsp),
	)

	service.Init()

	sayClent := model.NewService("tuyoo.micro.srv.auth", service.Client())
	fmt.Println(sayClent)
	rsp, err := sayClent.MakeAccessToken(context.Background(), &model.Request{
		UserName: "wuxun",
		UserId: 123,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp)
	//if err := service.Run(); err != nil {
	//	log.Fatal(err)
	//}
}