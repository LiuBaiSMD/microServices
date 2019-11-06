//此包用于实现具体web处理方法

package handler

import (
	"github.com/micro/go-micro/util/log"
	"handlerManageTest/myregistry"
	"net/http"
)

type RfAddr struct {}

func Init(){
	var rfaddr RfAddr
	myregistry.Registery(&rfaddr)
}

func (b* RfAddr)TestUserLogin() myregistry.HttpWR{
	return func(w http.ResponseWriter, r *http.Request){
		log.Log("method:", r.Method) //获取请求的方法
		log.Log("handlerfunc TestUserLogin")
	}
}

func (b* RfAddr) Login() myregistry.HttpWR {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Log("method:", r.Method) //获取请求的方法
		log.Log("handlerfunc Login")
	}
}
