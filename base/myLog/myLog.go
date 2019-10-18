package myLog

import (
	"fmt"
	"log"
	"time"
	"net/http"
)

type MyLogger interface {
	Init() error
	Info(info ...interface{})
	Debug(info ...interface{})
	Heart(info ...interface{})
	Error(info ...interface{})
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
//return
	ifInited = true
	dftLogger.Init()
	go heartBeat()
	handler = make(map[string]func(info ...interface{})error)
	handler["report"] = report
}

func NewLogger() MyLogger{
	return newLogger()
}

func newLogger() MyLogger{
	Logger = &dftLogger
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
