package util

import (
	"net/http"
	"errors"
	"io/ioutil"
	"encoding/json"
)

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