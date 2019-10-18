package myLog

import (
	"fmt"
	"log"
	"time"
	"net/http"
	"os"
)

type MyLogger interface {
	Init() error
	Info(info ...interface{})
	Debug(info ...interface{})
	Heart(info ...interface{})
	Error(info ...interface{})
	LogWithFile(fileName, prefix string, info ...interface{})
}
var (
	defaultFileName  = "log.log"
	dftLogger  = defaultLogger{}
	logger *log.Logger
	Logger MyLogger
	ifInited = false
	reportURL = "http://localhost:8080/userlogin"
)

var handler map[string]func(info ...interface{})error


func init()  {
	if !ifInited {
		dftLogger.Init()
		//启动心跳
		go heartBeat()
		handler = make(map[string]func(info ...interface{})error)
		handler["report"] = report
		Logger = &dftLogger
		ifInited = true
	}
}

func GetLogger() MyLogger{
	return Logger
}

func heartBeat(){
	tick := time.Tick(time.Second * 10)
	for _= range tick {
		Logger.Heart("pong pong pong ...")
		//Logger.Error("我要死了，救我！！！")
	}
}

func SetReport(report func(info ...interface{})error){
	handler["report"] = report
}

func SetLogPath(logPath string){
	//提供外部接口设置log的位置,如不设置则默认存放在脚本执行目录，名为log.log
	defaultFileName = logPath
}

func report(info ...interface{})error{
	rsp, err := http.Get(reportURL)
	fmt.Println(rsp, err, "   info: ", info)
	if err != nil{
		return err
	}
	return nil
}

func openLogFile(fileName string) (*os.File, error){
	if _,err :=os.Open(fileName);err!=nil && os.IsNotExist(err){
		os.Create(fileName)
	}
	logFile,err  := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND,0)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}