package crontabTask

import (
	"Crontab/task-dao"
	"encoding/json"
	"net/http"
	"reflect"
	"runtime"
	"time"
)

var taskFunc = make(map[string]func(...interface{})([]interface{}, error))
//var TaskKey = "TasksNow"
var client = taskdao.GetRedis()
type TaskInfo struct{
	Params []interface{}
	FuncName string
}

//将使用的方法进行添加注册
func AddTask(taskF func(...interface{})([]interface{}, error), params ...interface{}){

	fName := runtime.FuncForPC(reflect.ValueOf(taskF).Pointer()).Name()
	taskFunc[fName] = taskF

}

//设定interval，注册程序，并开始执行程序
func TimeRun(TaskKey string, f func(...interface{})([]interface{}, error), timeSecond int,params ...interface{}){
	fName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	taskFunc[fName] = f
	taskinfo := TaskInfo{
		Params: params,
		FuncName:fName,
	}
	taskinfoDump, _ := json.Marshal(taskinfo)
	tick := time.NewTicker(time.Duration(timeSecond) * time.Second)
	defer tick.Stop()
	for {
		client.LPush(TaskKey, taskinfoDump)
		<-tick.C
	}
}

//从redis队列中取任务
func RunTaskNow(TaskKey string, ){
	var taskinfo TaskInfo
	taskF, _ := client.RPop(TaskKey).Result()
	json.Unmarshal([]byte(string(taskF)), &taskinfo)
	f, ifOk := taskFunc[taskinfo.FuncName]
	if ifOk{
		go f(taskinfo.Params...)
	}
}




//params[0] ipAddressit
func CheckIfConnect(params ...interface{}){
	ipAddress := params[0]
	resp, err := http.Get(ipAddress.(string))
	if err!= nil{
		return
	}
	if resp.StatusCode ==  200 {
		return
	}
}