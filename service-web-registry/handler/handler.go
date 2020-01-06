//此包用于实现具体web处理方法

package handler

import (
	"github.com/micro/go-micro/util/log"
	"github.com/LiuBaiSMD/microServices/util/myregistry"
	"net/http"
	"errors"
	"io/ioutil"
	"encoding/json"
)

type RfAddr struct {}

func Init(){
	var rfaddr RfAddr
	myregistry.Registery(&rfaddr)
}

func (b* RfAddr)TestUserLogin() myregistry.HttpWR{
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		body, _ := GetBody(r)
		log.Log("method:", r.Method, "\nbody", body) //获取请求的方法
	}
}

func (b* RfAddr) Login() myregistry.HttpWR {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		body, _ := GetBody(r)
		log.Log("method:", r.Method, "\nbody:  ", body) //获取请求的方法		log.Log("handlerfunc Login")
	}
}
func GetBody(r *http.Request) (map[string]interface{}, error){
	//将参数解析为 map[string]interface{}型
	if r.Method=="GET"{
		r.ParseForm()
		res := make(map[string]interface{})
		for k, v := range r.Form{
			res[k] = v
		}
		return res, nil
	}
	if r.Method != "POST"{
		return nil, errors.New("请求类型错误，请检查")
	}
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
		if err := json.Unmarshal(b, &webData); err!=nil{
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