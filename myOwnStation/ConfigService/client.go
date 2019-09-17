package main

import (
	"fmt"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/consul"
	"github.com/micro/go-micro/util/log"
)

func main() {
	// 开始测试使用本地配置
	//appPath, _ := filepath.Abs(filepath.Dir(filepath.Join("./", string(filepath.Separator))))
	// 注册consul的配置地址
	consulSource := consul.NewSource(
		consul.WithAddress("127.0.0.1:8500"),
		consul.WithPrefix("/micro/config"),
		// optionally strip the provided prefix from the keys, defaults to false
		consul.StripPrefix(true),
	)
	// 创建新的配置
	conf := config.NewConfig()
	if err := conf.Load(consulSource); err != nil {
		log.Logf("load config errr!", err)
	}
	conf.Get("cluster", "consul")
	//if err := conf.Get("cluster", "consul"); err != nil {
	//	log.Logf("json format err!!!", err)
	//}
	//micro config的
	strMap := make(map[string]string)
	confData := conf.Get("cluster", "consul").StringMap(strMap)
	fmt.Println(confData)
	configMap := conf.Map()
	//fmt.Println(conf)
	fmt.Println(configMap)
	//fmt.Println(configMap[0])
}
