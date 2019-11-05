//此包用于实现具体web处理方法

package handler

import (
	"github.com/micro/go-micro/util/log"
	"handlerManageTest/myregistry"
	"net/http"
	"reflect"
)
type ControllerMapsType map[string]reflect.Value

type RfAddr struct {}

type Base struct{
	crMap ControllerMapsType
	funcRegistry map[string] myregistry.HttpWR
}

func Init(){
	var rfaddr RfAddr
	myregistry.Register.Registery(&rfaddr)
}


func (b* RfAddr)TestUserLogin() myregistry.HttpWR{
	f := func(w http.ResponseWriter, r *http.Request){
		log.Log("method:", r.Method) //获取请求的方法
		log.Log("handlerfunc TestUserLogin")
	}
	return f
}

func (b* RfAddr) Login() myregistry.HttpWR {
	f := func(w http.ResponseWriter, r *http.Request) {
		log.Log("method:", r.Method) //获取请求的方法
		log.Log("handlerfunc Login")
	}
	return f
}
