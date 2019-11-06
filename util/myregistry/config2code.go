//此模块用于读取配置，将url与handle绑定

package myregistry

import (
	baseJson "encoding/json"
	"errors"
	"fmt"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/encoder/json"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/config/source/file"
	"os"
	"path"
)

var headStr = `//自动生成的接口文件，package名字有误请自行更改
package %s

import (
	"net/http"
)

type RfAddr struct {}

func Init(){
	var rfaddr RfAddr
	myregistry.Registery(&rfaddr)
}
`

var handleStr = `
func (b* RfAddr)%s() myregistry.HttpWR{
	return func(w http.ResponseWriter, r *http.Request){
		// your handle logic todo ...
	}
}`

type rules struct{
	Func string `json:"Func"`
	Url string `json:"Url"`
}

func CodeFactory(configPath, dst string){
	dir := path.Dir(dst)
	if err := CheckDirOrCreate(dir);err!=nil{
		panic(err)
	}
	pkgName := path.Base(dir)
	var thisHeader = fmt.Sprintf(headStr, pkgName)
	conf, _ := ReadConfig(configPath)
	createPath := dst
	ifExist, _ := PathExists(createPath)
	var f *os.File
	f, _ = os.OpenFile(createPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if !ifExist{
		fmt.Println("没有我来造")
		if _, err :=f.Write([]byte(thisHeader)); err !=nil{
			panic(err)
		}
	}
	defer f.Close()
	for _,v := range conf{
		var r rules
		bv, _ := baseJson.Marshal(v)
		if err := baseJson.Unmarshal(bv, &r); err != nil{
			panic(err)
		}
		app := fmt.Sprintf(handleStr, r.Func)
		if _, err:= f.Write([]byte(app)) ;err != nil{
			panic(err)
		}
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

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//检查一个dir路径，没有则会创建
func CheckDirOrCreate(dirPath string) error{
	if ifExist,err :=PathExists(dirPath); err != nil{
		return err
	}else if !ifExist{
		err1 := os.MkdirAll(dirPath, 0777)
		if err1!=nil{
			return err1
		}
	}
	return nil
}