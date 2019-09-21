package dao

import (
	"github.com/go-redis/redis"
	"github.com/micro/go-micro/util/log"
	"github.com/LiuBaiSMD/microServices/base/config"
	"sync"
)

var (
	client *redis.Client
	m      sync.RWMutex
	inited bool
)

// Init 初始化Redis
func Init() {
	m.Lock()
	defer m.Unlock()

	if inited {
		log.Log("已经初始化过Redis...")
		return
	}
	config.Init()
	redisConfig := config.GetRedisConfig()
	// 打开才加载
	if redisConfig != nil {

		log.Log("初始化Redis...")
		log.Log("初始化Redis，普通模式...")
		initSingle(redisConfig)

		log.Log("初始化Redis，检测连接...")

		pong, err := client.Ping().Result()
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Log("初始化Redis，检测连接Ping.")
		log.Logf("初始化Redis，检测连接Ping... %s", pong)
	}
}

// GetRedis 获取redis
func GetRedis() *redis.Client {
	return client
}

func initSingle(redisConfig config.RedisConfig) {
	client = redis.NewClient(&redis.Options{
		Addr:     redisConfig.GetURL(),
		Password: redisConfig.GetPassword(), // no password set
		DB:       redisConfig.GetDB(),    // use default DB
	})
}

//func initSentinel(redisConfig config.RedisConfig) {
	//	client = redis.NewFailoverClient(&redis.FailoverOptions{
	//		MasterName:    redisConfig.GetSentinelConfig().GetMaster(),
	//		SentinelAddrs: redisConfig.GetSentinelConfig().GetNodes(),
	//		DB:            redisConfig.GetDBNum(),
	//		Password:      redisConfig.GetPassword(),
	//	})
	//
	//}