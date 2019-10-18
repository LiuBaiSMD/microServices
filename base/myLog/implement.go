package myLog

import (
	"log"
	"os"
	"runtime"
	"strings"
)

type defaultLogger struct {}

func (dfl *defaultLogger)Init() error{
	//初始化logger,别无他用
	if _,err :=os.Open(defaultFileName);err!=nil && os.IsNotExist(err){
		os.Create(defaultFileName)
	}

	logFile,err  := os.OpenFile(defaultFileName, os.O_RDWR|os.O_APPEND,0)
	//defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error !")
	}
	// 创建一个日志对象
	logger = log.New(logFile,"[Debug]",log.LstdFlags)
	//dftLogger := new(defaultLogger)
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

func getFileName()(string, int){
	_, file, line, _ := runtime.Caller(2)
	fileSplit := strings.Split(file, "/")
	fileName := fileSplit[len(fileSplit)-1]
	//fmt.Println(fileName)
	return fileName, line
}
