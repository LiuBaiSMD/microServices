package main

import (
	"GitWebhook/handler"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-micro/util/log"
	"os"
	"path"
)

func main() {
	service := web.NewService(
		web.Name("tuyoo.micro.web.tools"),
		web.Version("latest"),
		web.Address(":8010"),
		)
	log.Log("web.NewService")
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}
	log.Log("service.Init()")
	service.HandleFunc("/gitupdate", handler.GetPush)
	curPath,_ :=os.Getwd()
	gitBaseDir:=path.Join(curPath, "gits")
	handler.GetFilelist(gitBaseDir)
	log.Log("service.Run()")
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

