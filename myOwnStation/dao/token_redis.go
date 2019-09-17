package dao

import (
	"fmt"
	"github.com/go-redis/redis"
	"errors"
	"strconv"
	"time"
	"github.com/micro/go-micro/util/log"
)

var rdsConn *redis.Client

func InitTokenRedis(connType, redisUrl string) *redis.Client {
	rdsConn = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := rdsConn.Ping().Result()
	if err != nil{
		fmt.Println(pong, err)
		return nil
	}
	// Output: PONG <nil>
	return rdsConn
}

func SetRediString(key , value string, expire time.Duration){
	log.Logf("插入数据：%v  :%v", key, value)
	err := rdsConn.Set(key, value, expire).Err()
	if err != nil {
		panic(err)
	}
}

func CreateUserIdToken(userId string) string{
	token := userId
	stampStr :=strconv.FormatInt(time.Now().Unix(), 10)
	token = token + "_" + stampStr
	return token
}

func CheckToken(userId, rqsToken string) error{
	//检查token值是否存在，并返回成功标志
	rdsToken := rdsConn.Get(userId)
	if rdsToken.Val() != rqsToken{
		return errors.New("无效Token!")
	}
	return nil
}