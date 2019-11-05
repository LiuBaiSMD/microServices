package handler

import (
	httpJson "encoding/json"
	"errors"
	"fmt"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/encoder/json"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/config/source/file"
	"handlerManageTest/myregistry"
	"io/ioutil"
	"net/http"
	"github.com/micro/go-micro/web"
)

func HandlerFromConf(service web.Service, configPath string){
	conf, _ := ReadConfig(configPath)
	for _,v := range conf{
		fName, _ := GetMapContent(v.(map[string]interface{}), "func")
		url, _ := GetMapContent(v.(map[string]interface{}), "url")
		fmt.Println("f ------> url  :  ", fName, url)
		myregistry.BindUrlHandle(service, url.(string), fName.(string))
	}
}

func GetBody(r *http.Request) (map[string]interface{}, error){
	//将参数解析为 map[string]interface{}型
	ContType  := r.Header["Content-Type"]
	if ContType[0] == "application/json"{
		if err:=r.ParseForm();err!=nil{
			return nil, errors.New("参数解析异常")
		}
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, errors.New("连接错误")
		}
		var webData interface{}
		if err := httpJson.Unmarshal(b, &webData); err!=nil{
			return nil, errors.New("json解析异常")
		}
		mapdata := webData.(map[string]interface{})
		return mapdata, nil
	}
	if ContType[0] == "application/x-www-form-urlencoded"{
		r.ParseForm()
		var mapdata map[string]interface{}
		mapdata = make(map[string]interface{})
		for k,v := range r.Form{
			mapdata[k] = v[0]
		}
		return mapdata, nil
	}
	return nil, errors.New("请求HEADER类型错误，请检查！")
}

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