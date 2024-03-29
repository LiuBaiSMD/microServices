package myLog

import (
	"log"
	"runtime"
	"strings"
)

type defaultLogger struct {}

func (dfl *defaultLogger)Init() error{
	//初始化logger,别无他用
	logFile,err := openLogFile(defaultFileName)
	if err != nil{
		log.Fatalf("日志文件打开失败！", defaultFileName)
	}
	// 创建一个日志对象
	logger = log.New(logFile,"[Debug]",log.LstdFlags)
	return nil
}

func (dfl *defaultLogger)Info(info ...interface{}){
	fileName, line := getFileName()
	logger.SetPrefix("[Info]")
	logger.Println(fileName, line, info)
}

func (dfl *defaultLogger)Debug(info ...interface{}){
	fileName, line := getFileName()
	logger.SetPrefix("[Debug]")
	logger.Println(fileName, line, info)
}

func (dfl *defaultLogger)Heart(info ...interface{}){
	fileName, line := getFileName()
	logger.SetPrefix("[Heart]")
	logger.Println(fileName, line, info)
}

func (dfl *defaultLogger)Error(info ...interface{}){
	fileName, line := getFileName()
	//防止report出现错误
	err := handler["report"](info...)
	if err != nil{
		logger.SetPrefix("[Error Report]")
		logger.Println(err)
	}
	logger.SetPrefix("[Error]")
	logger.Println(fileName, line, info)
}

func (dfl *defaultLogger)LogWithFile(fileName, prefix string, info ...interface{}){
	logFile, err := openLogFile(fileName)
	if err != nil{
		log.Fatalf("日志文件打开失败: ",err.Error())
	}
	defer logFile.Close()
	// 创建一个日志对象
	fLogger := log.New(logFile,"[Debug]",log.LstdFlags)
	debugFile, line := getFileName()
	fLogger.SetPrefix(prefix)
	fLogger.Println(debugFile, line, info)
}

func getFileName()(string, int){
	_, file, line, _ := runtime.Caller(2)
	fileSplit := strings.Split(file, "/")
	fileName := fileSplit[len(fileSplit)-1]
	return fileName, line
}


