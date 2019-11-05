//此模块用于读取配置，将url与handle绑定

package handler

import (
	"fmt"
	"github.com/micro/go-micro/web"
	"handlerManageTest/myregistry"
)

//读取配置，将方法与路由绑定在web.Service上
func BindHandlerFromConf(service web.Service, configPath string){
	conf, _ := ReadConfig(configPath)
	for _,v := range conf{
		fName, _ := GetMapContent(v.(map[string]interface{}), "func")
		url, _ := GetMapContent(v.(map[string]interface{}), "url")
		fmt.Println("f ------> url  :  ", fName, url)
		myregistry.BindUrlHandle(service, url.(string), fName.(string))
	}
}
