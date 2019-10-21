## myLog模块使用简介
### 1.自动化部分
```markdown
1.在导入包后会自动启动日志心跳输出
2.文件名根据project.json文件logPath + name + date.log(今日日期2016-1-1) 创建日志文件
3.如文件已存在，则会追加写入

```

2.使用模板
需要在main.go同级别目录中存在project.json文件，
```markdown
// project.json项目配置文件

{
  "logPath": "../logs/",
  "name": "microLog"
}
```
```markdown
// example.go使用示范

package main

import (
	"fmt"
	"github.com/LiuBaiSMD/microServices/base/myLog"
	"net/http"
)

func main(){
	myLog.SetReport(report) //设置自定义的错误上报函数 func report(info ...interface{})error
	logger := myLog.GetLogger() //使用GetLogger()获取 logger
	logger.Debug("test123")
	logger.Error("我来看看有没有错误!")
	logger.LogWithFile("test.log", "hhh")
	for{
		continue
	}
}

func report(info ...interface{})error{
    //report方法规则是将传入的信息拼接在url后直接转发
	infoStr := fmt.Sprint(info...)
	url := "http://localhost:8080/errorReport?info=" + infoStr
	_, err := http.Get(url)
	if err != nil{
		return err
	}
	return nil
}
```