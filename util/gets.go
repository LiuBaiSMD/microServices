package util

import (
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-micro/config/encoder/json"
	"reflect"
	"errors"
	"time"
	"strings"
	"runtime"
)
func GetType(c interface{}) string{
	return reflect.TypeOf(c).String()
}

func GetTypes(c ...interface{}) []string{
	//
	var res []string
	for _, k := range c{
		t := reflect.TypeOf(k).String()
		res = append(res, t)
	}
	return res
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

func GetConfig(filePath string)(map[string]interface{}, error){
	//从一个文件中读取配置，
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

func GetToday()string{
	data := time.Now().Format("2006-01-02")
	return data
}

func GetUseFileWithLine(skip int)(string, int){
	//获取调用此文件的文件名, 行数，通过skip设定层级
	_, file, line, _ := runtime.Caller(2)
	fileSplit := strings.Split(file, "/")
	fileName := fileSplit[len(fileSplit)-1]
	return fileName, line
}