package handler

import (
	"errors"
	"fmt"
	"github.com/LiuBaiSMD/microServices/base/config"
	"github.com/LiuBaiSMD/microServices/util"
	"github.com/micro/go-micro/util/log"
	"net/http"
	"service-discuss-room/dao"
	"time"
)

var inited  = false

type Auth struct {
	Id       string `gorm:"default:'peter'"`
	Password string
	Name	 string
}

func Init(){
	if !inited{
		log.Log("初始化handler模块！")
		config.Init()
		dao.Init(
			dao.SetRedisPassword(config.GetRedisConfig().RedisPassword),
			dao.SetRedisUrl(config.GetRedisConfig().RedisUrl),
			dao.SetMysqlDriveName(config.GetMysqlConfig().MysqlDriveName),
			dao.SetMysqlURL(config.GetMysqlConfig().MysqlURL),
			)
		InitDiscuss()
		//config.Init()
		util.Init()
		inited =true
		//log.Log("MyOwnStation config:	", config.MyOwnStation)
	}
}

func SetDiscussReq(w http.ResponseWriter, r *http.Request){
	//defer r.Body.Close()
	//log.Log("SetDiscuss	! methos:	", r.Method)
	if !util.CheckReqAllowed(r.Method, "POST") {
		fmt.Fprintln(w, "不接受的请求类型:	" + r.Method)
		log.Log("不接受的请求类型！" + r.Method)
		return
	}
	mapdata, err := util.GetBody(r)
	if err != nil{
		log.Log("GetBody参数解析错误！")
		fmt.Fprint(w,err)
	}
	chatRoom, ok1 := mapdata["chatRoom"].(string)
	userId, ok2 := mapdata["userId"].(string)
	content, ok3 := mapdata["content"].(string)
	if !util.CheckOKs(ok1, ok2, ok3){
		fmt.Fprint(w, errors.New("参数解析失败!"))
		log.Log("参数解析失败!")
	}
	nowStamp := float64(time.Now().Unix())
	if err := setDiscuss(chatRoom, userId, content, nowStamp);err!=nil{
		log.Log("存储评论错误：	", err)
		fmt.Fprint(w, err)
	}
	fmt.Fprint(w, "操作成功!")
	//cont := Content{
	//	Context:content,
	//}
	//stampSTR := strconv.FormatInt(time.Now().Unix(), 10)
	//member := Discuss{
	//	UserId: userId,
	//	DiscussKey: stampSTR,
	//}
	//mashMember, err := json.Marshal(member)
	//if err != nil{
	//	fmt.Fprint(w, errors.New("存储member序列化失败！"))
	//	return
	//}
	//
	//added, err := DiscussZAdd(chatRoom, float64(time.Now().Unix()), mashMember)
	//if err != nil{
	//	fmt.Fprint(w, err)
	//	return
	//}

	log.Log("成功插入!")
}

func GetDiscussReq(w http.ResponseWriter, r *http.Request){
	//log.Log("GetDiscuss	! methos:	", r.Method)
	if r.Method != "POST" {
		fmt.Fprintln(w, "只接受post请求")
		log.Log("只接受Post请求！")
		return
	}
	mapdata, err := util.GetBody(r)
	if err != nil {
		log.Log("GetBody参数解析错误！")
		fmt.Fprint(w,err)
	}
	//start
	chatRoom, ok1 := mapdata["chatRoom"].(string)  //strconv.FormatInt
	start := int64(mapdata["start"].(float64))
	stop := int64(mapdata["stop"].(float64))//mapdata["stop"]
	//log.Log("oks:	", ok2, ok3)
	//log.Logf("%T %T", start, stop)
	//log.Log(chatRoom, start, stop)
	if !util.CheckOKs(ok1){
		//log.Log(ok1,)
		fmt.Fprint(w, errors.New("参数解析失败!"))
		log.Log("参数解析失败!", mapdata, chatRoom, start, stop)
	}
	res, err := DiscussZRevRangeWithScores(chatRoom, start, stop)
	if err != nil{
		fmt.Fprint(w, err)
		return
	}
	//for _, v :=range res{
	//	fmt.Println(v.Time)
	//	//fmt.Println(v.UserId, ": ", v.Content, "\n")
	//}
	//log.Log(res)
	//for k, v := range res{
	//	log.Log(k, ": ", v)
	//}
	fmt.Fprint(w, res)
}

func DiscussOtherReq(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintln(w, "只接受post请求")
		log.Log("只接受Post请求！")
		return
	}
	mapdata, err := util.GetBody(r)
	if err != nil {
		log.Log("GetBody参数解析错误！")
		fmt.Fprint(w,err)
	}
	preUserId, _ := mapdata["preUserId"].(string)
	preTime, _ := mapdata["preTime"].(string)
	userId, _ := mapdata["userId"].(string)
	chatRoom, _ := mapdata["chatRoom"].(string)
	context, _ := mapdata["context"].(string)
	if err1 := discussOtherZAdd(preUserId, preTime, userId, chatRoom, context);err1 != nil{
		fmt.Fprint(w, err1)
	}
	fmt.Fprint(w, "操作成功！")
}