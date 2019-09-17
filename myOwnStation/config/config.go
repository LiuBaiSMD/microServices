package config

import (
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/consul"
	"github.com/micro/go-micro/util/log"
	"os"
	"sync"
	"time"
)

var (
	err error
)

type defaultMyOwnStation struct {
	MyOwnStation        string `json:"myownstation"`
	TokenTTL        	int    `json:"token_ttl"`
}

var (
	defaultConfigPath       = "myOwnStation" // 默认的仓库地址
	defaultConsulServerAddr = "localhost:8500"
	dockerConsulServerAddr = "consul4:8500"
	m                       sync.RWMutex
	inited                  bool
	MyOwnStation 			defaultMyOwnStation
	webConf					webConfig
	consulConf				consulConfig
	redisConf				redisConfig
	mysqlConf				mysqlConfig
)

// Init 初始化配置
func Init() {

	m.Lock()
	defer m.Unlock()

	if inited {
		log.Logf("[Init] 配置已经初始化过")
		return
	}
	dockerMode := os.Getenv("RUN_DOCKER_MODE")
	if dockerMode == "on" {
		log.Log("docker模式")
		defaultConsulServerAddr = dockerConsulServerAddr
	}

	// 从注册中心读取配置
	consulSource := consul.NewSource(
		consul.WithAddress(defaultConsulServerAddr),
		consul.WithPrefix(defaultConfigPath),
		consul.StripPrefix(true),
	)
	// 创建新的配置
	conf := config.NewConfig()
	for{
		if err := conf.Load(consulSource); err != nil {
			log.Logf("load config errr!!!", err)
			time.Sleep(time.Second * 5)
			continue
		}
		break
	}

	// 侦听文件变动
	watcher, err := conf.Watch()
	if err != nil {
		log.Fatalf("[Init] 侦听consul配置中心 watcher异常，%s", err)
		panic(err)
	}
	go func() {
		for {
			v, err := watcher.Next()
			if err != nil {
				log.Fatalf("侦听consul配置中心 异常， %s", err)
				return
			}
			if err = conf.Load(consulSource); err != nil {
				panic(err)
			}
			log.Logf("consul配置中心有变化，%s", string(v.Bytes()))
			readConfig(conf)
			log.Log(MyOwnStation, webConf, consulConf, redisConf, mysqlConf)
		}
	}()
	// 赋值
	readConfig(conf)
	log.Log(MyOwnStation, webConf, consulConf, redisConf, mysqlConf)
	inited = true
	// 标记已经初始化
}

func readConfig(conf config.Config) error{
	if err := conf.Get("config").Scan(&MyOwnStation); err != nil {
		log.Logf("总配置加载异常:%s", err)
		return err
	}
	if err := conf.Get("config", "web_config").Scan(&webConf); err != nil {
		log.Logf("webService配置加载异常:%s", err)
		return err
	}
	if err := conf.Get("config", "consul_config").Scan(&consulConf); err != nil {
		log.Logf("总配置加载异常:%s", err)
		return err
	}
	if err := conf.Get("config", "redis_config").Scan(&redisConf); err != nil {
		log.Logf("总配置加载异常:%s", err)
		return err
	}
	if err := conf.Get("config", "mysql_config").Scan(&mysqlConf); err != nil {
		log.Logf("总配置加载异常:%s", err)
		return err
	}
	return nil
}

func GetWebConfig() webConfig{
	log.Log("getWebConfig:	", webConf)
	return webConf
}

func GetConsulConfig() consulConfig{
	log.Log("getonsulConfig:	", consulConf)
	return consulConf
}

func GetRedisConfig() redisConfig{
	log.Log("getredisConfig:	", redisConf)
	return redisConf
}

func GetMysqlConfig() mysqlConfig{
	log.Log("getmysqlConfig:	", mysqlConf)
	return mysqlConf
}





