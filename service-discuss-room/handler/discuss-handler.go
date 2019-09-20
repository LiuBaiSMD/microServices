package handler

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/micro/go-micro/util/log"
	"github.com/LiuBaiSMD/microServices/config"
	
	"github.com/LiuBaiSMD/microServices/dao" //
	"github.com/LiuBaiSMD/microServices/util"
	"time"
	"strconv"
	"errors"
)

var (
	initDiscuss = false
	discussRdsConn *redis.Client
	err error
	)

type Content struct {
	Context	string	`json:"content"`
	SonDiscuss	[]Discuss
}

type Discuss struct {
	UserId		string	`json:"userid"`
	DiscussKey		string	`json:"discuss_key"` //Content
}

type DiscussResult struct {
	UserId		string	`json:"userid"`
	Context		string	`json:"discuss_key"` //Content
	Time	float64
}

type ContentResult struct {
	UserId		string	`json:"userid"`
	DiscussKey		string	`json:"discuss_key"` //Content
	Time	float64
	SonDiscuss	[]Discuss
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
		for i:=0; i<1; i++{
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

func setDiscuss(chatRoom, userId , context string, nowStamp float64)error{
	stampSTR := strconv.FormatInt(time.Now().Unix(), 10)
	cont := Content{
		Context:context,
	}
	member := Discuss{
		UserId: userId,
		DiscussKey: stampSTR,
	}
	mashMember, err := json.Marshal(member)
	if err != nil{
		log.Log("存储member序列化失败！")
		return errors.New("存储member序列化失败！")
	}
	_, err1 := discussZAdd(chatRoom, nowStamp, mashMember) //float64(time.Now().Unix())
	if err1 != nil{
		log.Log(err1)
		return err1
	}
	contMar, _ := json.Marshal(cont)
	//存入专门的hash库
	discussRdsConn.HSet(getChatRoomText(chatRoom), userId + ":" + stampSTR, contMar)
	return nil
}

func discussZAdd(chatRoom string, score float64, Member interface{})(int64, error){
	//log.Log(chatRoom, score, Member)
	z := redis.Z{
		Score:score,
		Member:Member,
	}
	added, err := discussRdsConn.ZAdd(chatRoom, z).Result()
	return added, err
}

func DiscussZRevRangeWithScores(chatRoom string, start , stop int64)([]DiscussResult, error){
	result, err := discussRdsConn.ZRevRangeWithScores(chatRoom, start, stop).Result()
	if err != nil{
		return nil,err
	}
	var dress []DiscussResult
	for _, v := range result{
		var resultDiscuss Discuss
		var contentResult Content
		var dres		DiscussResult
		json.Unmarshal([]byte(v.Member.(string)), &resultDiscuss)
		tt := strconv.FormatFloat(v.Score, 'f', -1, 64)
		DiscussKey := resultDiscuss.UserId + ":" + tt
		dis, err := discussRdsConn.HGet(getChatRoomText(chatRoom), DiscussKey).Result()
		if err !=nil{
			return nil, errors.New("读取评论错误！")
		}
		//fmt.Println("timeint:	", tt)
		json.Unmarshal([]byte(dis), &contentResult)
		dres.UserId = resultDiscuss.UserId
		dres.Time = v.Score
		dres.Context = contentResult.Context
		dress = append(dress, dres)

	}
	return dress, nil
}

func getChatRoomText(chatRoom string) string{
	return chatRoom + "text"
}

func discussOtherZAdd(preUserId, preTime, userId, chatRoom, context string)error{
	nowStamp := time.Now().Unix()
	if err:= setDiscuss(chatRoom, userId, context, float64(nowStamp));err!=nil{
		return err
	}
	DiscussKey := preUserId + ":" + preTime
	//log.Log("chatroomText:	", getChatRoomText(chatRoom), DiscussKey)
	dis, err := discussRdsConn.HGet(getChatRoomText(chatRoom), DiscussKey).Result()
	//log.Log("dis:	", dis)
	if err !=nil{
		return errors.New("读取评论错误！")
	}
	var contentResult Content
	json.Unmarshal([]byte(dis), &contentResult)
	nowStampSTR := strconv.FormatInt(nowStamp, 10)
	contentResult.SonDiscuss = append(contentResult.SonDiscuss, Discuss{
		UserId:userId,
		DiscussKey:nowStampSTR,
	})
	contentResultMash,_ := json.Marshal(contentResult)
	ifOK, err := discussRdsConn.HSet(getChatRoomText(chatRoom), preUserId + ":" + preTime, contentResultMash).Result()
	if !ifOK || err !=nil{
		return err
	}
	return nil
}
