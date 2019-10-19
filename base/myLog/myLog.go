package myLog

import (
	"github.com/LiuBaiSMD/microServices/util"
	"log"
	"time"
	"net/http"
	"os"
	"strings"
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
	projectConf = "project.json"
	defaultFileName  string
	dftLogger  = defaultLogger{}
	logger *log.Logger
	Logger MyLogger
	ifInited = false
	reportURL = "http://localhost:8080/userlogin"
)

var handler map[string]func(info ...interface{})error


func init()  {
	if !ifInited {
		defaultFileName = getDefaultLogFileName()
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
	log.Println(rsp, err, "   info: ", info)
	if err != nil{
		return err
	}
	return nil
}

func openLogFile(fileName string) (*os.File, error){
	s := strings.Split(fileName,"/")
	file := s[len(s)-1]
	dirPath := fileName[0:len(fileName)-len(file)]
	if dirPath != "" { //判断是否存在dirPath
		err := os.MkdirAll(dirPath , os.ModePerm)
		if err!=nil{
			return nil, err
		}
	}
	logFile,err  := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE,0666)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}

func GetToday()string{
	data := time.Now().Format("2006-01-02")
	return data
}

func getDefaultLogFileName()string{
	if ifExist, err := util.CheckPathExists(projectConf);(!ifExist || err !=nil){
		log.Println("do no find file -----> project.json\ndefault log file ------> yourLog.log ")
		return "yourLog.log"
	}else {
		cfg, err := util.GetConfig(projectConf)
		if err != nil{
			panic(err)
		}
		name, _ := util.GetMapContent(cfg, "name")
		logPath, _ := util.GetMapContent(cfg, "logPath")
		data := util.GetToday()
		return logPath.(string) + "/" + name.(string) + "/MS-Log-" + data + ".log"
	}

}