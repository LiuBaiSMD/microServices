//此模块用于读取配置，将url与handle绑定

package myregistry

import (
	baseJson "encoding/json"
	"fmt"
	"github.com/LiuBaiSMD/microServices/util"
	"os"
	"path"
)

var headStr = `//自动生成的接口文件，package名字有误请自行更改
package %s

import (
	"net/http"
	"github.com/LiuBaiSMD/microServices/util/myregistry"
)


//>>>>>>>>>>如第一次使用请取消注释下方代码<<<<<<<<<<<<

//type RfAddr struct {}
//
//func Init(){
//	var rfaddr RfAddr
//	myregistry.Registery(&rfaddr)
//}
`

var handleStr = `
func (b* RfAddr)%s() myregistry.HttpWR{
	return func(w http.ResponseWriter, r *http.Request){
		// your handle logic todo ...
	}
}
`

type rules struct{
	Func string `json:"Func"`
	Url string `json:"Url"`
}

//dst请配置绝对路径
func CodeFactory(configPath, dst string){
	dir := path.Dir(dst)
	if err := util.CheckDirOrCreate(dir);err!=nil{
		panic(err)
	}
	pkgName := path.Base(dir)
	var thisHeader = fmt.Sprintf(headStr, pkgName)
	conf, _ := util.GetConfig(configPath)
	createPath := dst
	ifExist, _ := util.CheckPathExists(createPath)
	var f *os.File
	f, _ = os.OpenFile(createPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if !ifExist{
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
