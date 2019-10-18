package handler


import (
	"github.com/micro/go-micro/web"
	"net/http"
	"fmt"
)

type handlerFuncUrl map[string]func(w http.ResponseWriter, r *http.Request)  //配置handlerFunc表的数据结构
type handlerUrl map[string]http.Handler

var WebConfig = handlerFuncUrl{
	"/userlogin/": UserLogin,
	"/tokenLogin": TokenLogin,
	"/userregister/": Register,
	"/changePWD": ChangePWDReq,
	"/errorReport": ErrorReport,
}

func SetHandleFunc(service web.Service, WebConfig handlerFuncUrl){
	for k, v := range WebConfig{
		service.HandleFunc(k, v)
	}
}

var WebHandlerConfig = handlerUrl{
	"websocket": http.StripPrefix("/websocket/", http.FileServer(http.Dir("html/websocket"))),
	"changeTest":  http.StripPrefix("/changeTest/", http.FileServer(http.Dir("html/ChangeTest"))),
}

func SetHandle(service web.Service, WebHandlerConfig handlerUrl){
	for k, v := range WebHandlerConfig{
		service.Handle(k, v)
	}
}

func Report(info ...interface{})error{
	infoStr := fmt.Sprint(info...)
	url := "http://localhost:8080/errorReport?info=" + infoStr
	_, err := http.Get(url)
	if err != nil{
		return err
	}
	return nil
}