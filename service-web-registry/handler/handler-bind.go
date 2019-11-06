//此模块用于读取配置，将url与handle绑定

package handler

import (
	"fmt"
	"github.com/micro/go-micro/web"
	"handlerManageTest/myregistry"
	"encoding/json"
	//"github.com/goinggo/mapstructure"
)

type rules struct{
	Func string `json:"Func"`
	Url string `json:"Url"`
}
//读取配置，将方法与路由绑定在web.Service上
func BindHandlerFromConf(service web.Service, configPath string){
	conf, _ := ReadConfig(configPath)
	for _,v := range conf{
		var r rules
		//r := v.(rules)
		bv, _ := json.Marshal(v)
		json.Unmarshal(bv, &r)
		fmt.Println("url ------> handle  :  ", r.Url, r.Func)
		myregistry.BindUrlHandle(service, r.Url, r.Func)
	}
}
