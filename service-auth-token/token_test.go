/*
auth:   wuxun
date:   2020-01-13 15:15
mail:   lbwuxun@qq.com
desc:   how to use or use for what
*/


package main

import (
    "context"
    "fmt"
    "testing"
    micro "github.com/micro/go-micro"
    proto "JWTToken/proto" //这里写你的proto文件放置路劲
)


func TestJwtTokenCreator_GetToken(t *testing.T) {
    // Create a new service. Optionally include some options here.
    service := micro.NewService(micro.Name("go.micro.api.token_client"))
    //service.Init()

    // Create new greeter client
    greeter := proto.NewTokenCreatorService("go.micro.api.token", service.Client())

    // Call the greeter
    body := &proto.TokenRequest{
        Name:"wuxun",
        Uid:"123456",
    }
    rsp, err := greeter.GetToken(context.TODO(), body)
    if err != nil {
        fmt.Println(err)
    }

    // Print response
    fmt.Println(rsp)
}
