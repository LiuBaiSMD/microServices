/*
auth:   wuxun
date:   2020-01-15 10:21
mail:   lbwuxun@qq.com
desc:   how to use or use for what
*/

package main

import (
	"JWTToken/proto"
	"JWTToken/handler"
	"github.com/micro/go-micro"
	"log"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.token"),

	)
	service.Init()
	// register example handler
	jwtToken.RegisterTokenCreatorHandler(service.Server(), new(handler.JwtTokenCreator))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
