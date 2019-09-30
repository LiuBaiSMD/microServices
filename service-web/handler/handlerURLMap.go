package handler


import (
	"github.com/micro/go-micro/web"
	"net/http"
)

type handlerUrl map[string]func(w http.ResponseWriter, r *http.Request)

var WebConfig = handlerUrl{
	"/userlogin/": UserLogin,
	"/tokenLogin": TokenLogin,
	"/userregister/": Register,
	"/changePWD": ChangePWDReq,
}

func SetHandleFunc(service web.Service, WebConfig handlerUrl){
	for k, v := range WebConfig{
		service.HandleFunc(k, v)
	}
}