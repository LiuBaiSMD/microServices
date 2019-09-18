package handler

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/micro/go-micro/util/log"
	"microServices/config"
	"microServices/dao"
	"microServices/util"
	"time"
)

var (
	initDiscuss = false
	discussRdsConn *redis.Client
	err error
	)

type Discuss struct {
	UserId		string	`json:"userId"`
	Content		string	`json:"content"`
}

type discussResult struct {
	Score float64
	Member Discuss
}

func InitDiscuss(){
	if !initDiscuss{
		log.Log("初始化handler模块！")
		config.Init()
		util.Init()
		dao.Init(
			dao.SetRedisPassword(config.GetRedisConfig().RedisPassword),
			dao.SetRedisUrl(config.GetRedisConfig().RedisUrl),
			dao.SetMysqlDriveName(config.GetMysqlConfig().MysqlDriveName),
			dao.SetMysqlURL(config.GetMysqlConfig().MysqlURL),
		)
		c := ".."
		for i:=0; i<3; i++{
			log.Log("正在初始化redis."+ c)
			c = c + ".."
			time.Sleep(time.Second)
		}
		discussRdsConn, err = dao.GetRedisClient()
		if err != nil{
			log.Log(err)
			return
		}
		inited =true
	}
}

func DiscussZAdd(chatRoom string, score float64, Member interface{})(int64, error){
	log.Log(chatRoom, score, Member)
	z := redis.Z{
		Score:score,
		Member:Member,
	}
	added, err := discussRdsConn.ZAdd(chatRoom, z).Result()
	return added, err
}

func DiscussZRevRangeWithScores(chatRoom string, start , stop int)([]Discuss, error){
	result, err := discussRdsConn.ZRevRangeWithScores(chatRoom, 0, -1).Result()
	if err != nil{
		return nil,err
	}
	var res []Discuss
	for _, v := range result{
		var result1 Discuss
		json.Unmarshal([]byte(v.Member.(string)), &result1)
		res = append(res, result1)
	}
	return res, nil
}
