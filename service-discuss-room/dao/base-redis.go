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

func InitRedis(Password, redisUrl string, DB int) *redis.Client { //InitTokenRedis
	rdsConn = redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: Password, // no password set
		DB:       DB,  // use default DB
	})
	rdsConn.BgRewriteAOF()
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

func GetRedisClient()(*redis.Client, error){
	if rdsConn!=nil{
		return rdsConn, nil
	}
	return nil, errors.New("redis连接失败！")
}