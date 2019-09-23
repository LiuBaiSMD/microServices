package main

import(
	"github.com/LiuBaiSMD/microServices/proto/auth"
	"github.com/micro/go-micro/client"
	"context"
	"fmt"
)

func main() {
	var authService auth.Service
	authService = auth.NewService("tuyoo.micro.srv.auth", client.DefaultClient)
	rspToken, err := authService.MakeAccessToken(context.TODO(), &auth.Request{
		UserName: "wuxun",
		UserId: 123,
	})
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(rspToken)
}