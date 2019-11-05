//此模块用于读取配置，将url与handle绑定

package handler

import (
	"errors"
	"fmt"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/encoder/json"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/config/source/file"
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

//获取字典中的值
func GetMapContent(m map[string]interface{}, path ...string) (interface{}, error){
	//本接口将获取一个map中，按path路径取值，返回一个interface
	var content interface{}
	var ok bool
	l := len(path)
	if l ==0 || (l == 1 && path[0]==""){  //当没有填入
		return m, nil
	}
	for k, v:= range path{
		if k ==l-1{
			content, ok = m[v]
			if !ok{
				return nil, errors.New(" 配置读取错误---> 	" + v)
			}
			return content,nil
		}
		if m, ok = m[v].(map[string]interface{}); !ok{
			return nil, errors.New(" 配置读取错误---> 	" + v)
		}
	}
	return nil, errors.New("missing map!")
}

//指定文件中的配置
func ReadConfig(filePath string)(map[string]interface{}, error){
	configPath := filePath
	e := json.NewEncoder()
	fileSource := file.NewSource(
		file.WithPath(configPath),
		source.WithEncoder(e),
	)
	conf := config.NewConfig()
	// 加载micro.yml文件
	if err := conf.Load(fileSource); err != nil {
		panic(err)
	}
	routes := make(map[string]interface{})
	err := conf.Scan(&routes)
	if err != nil{
		return nil, err
	}
	return routes, nil
}