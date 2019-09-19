package handler

import (
	"fmt"
	"github.com/micro/go-micro/util/log"
	"html/template"
	"microServices/config"
	"microServices/dao"
	"microServices/util"
	"net/http"
	"time"
	"errors"
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

func UserLogin(w http.ResponseWriter, r *http.Request) {
	log.Log("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		log.Logf("加载登录界面")
		t, _ := template.ParseFiles("html/userlogin.html")
		t.Execute(w, nil)
		log.Logf("加载完毕")
		return
	}
	t := template.New("test")
	//请求的是登陆数据，那么执行登陆的逻辑判断
	r.ParseForm()
	mapdata, err := util.GetBody(r)
	if err != nil{
		fmt.Fprintln(w, err.Error())
		return
	}

	name, ok1 := mapdata["name"].(string)
	password, ok2 := mapdata["password"].(string)
	ok := util.CheckOKs(ok1, ok2)
	if !ok {
		fmt.Fprintln(w, "params解析失败")
		return
	}
	if name == "" || password == "" {
		fmt.Fprintln(w, "请填写账号、密码！")
		return
	}
	if err := dao.QueryUserIdPass(name, password); err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	userToken := dao.CreateUserIdToken(name)
	dao.SetRediString(name, userToken, time.Second * 100 )
	t, _ = template.ParseFiles("html/websocket/index.html")
	log.Log("token:		", userToken)
	t.Execute(w, userToken)
}

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("html/register.html")
		t.Execute(w, "")

	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		mapdata, err := util.GetBody(r)
		if err != nil{
			fmt.Fprintln(w, err.Error())
			return
		}
		r.ParseForm()
		name, ok1 := mapdata["name"].(string)
		password, ok2 := mapdata["password"].(string)
		password2, ok3 := mapdata["password2"].(string)
		ok := util.CheckOKs(ok1, ok2, ok3)
		if !ok {
			fmt.Fprintln(w, "params解析失败!")
			return
		}
		//t := template.New("test")
		//进行参数检查
		if name == "" || password == "" || password2 == "" {
			fmt.Fprintln(w, "请完整填写账号、密码！")
			return
		}
		if name == "" {
			fmt.Fprintln(w, "请输入用户名！")
			return
		}
		if password == "" || password2 == "" {
			fmt.Fprintln(w, "密码不能为空")
			return
		} else if password != password2 {
			fmt.Fprintln(w, "两次密码不一致！请确认后重试！")
			return
		}
		//进行数据库检查
		if err := dao.RegisterUserIdPass(name, password); err != nil {
			fmt.Fprintln(w, err.Error())
			return
		}
		http.Redirect(w, r, "html/login.html", http.StatusOK)
	}
}


func ChangePWDReq(w http.ResponseWriter, r *http.Request) {
	log.Log("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		log.Logf("加载修改密码界面")
		t, _ := template.ParseFiles("html/changePWD/index.html")
		t.Execute(w, nil)
		log.Logf("加载完毕")
		return
		return
	} else{
		r.ParseForm()
		mapdata, err := util.GetBody(r)
		if err != nil {
			fmt.Fprintln(w, err.Error())
			return
		}
		name, ok1 := mapdata["userId"].(string)
		password, ok2 := mapdata["password"].(string)
		newPassword, ok3 := mapdata["newPassword"].(string)
		newPassword2, ok4 := mapdata["newPassword2"].(string)
		ok := util.CheckOKs(ok1, ok2, ok3, ok4)
		if !ok{
			fmt.Fprintln(w, "params解析失败")
			return
		}
		if name == "" || password == "" || newPassword == "" || newPassword2 == "" {
			fmt.Fprintln(w, "请完整填写账号、密码！")
			return
		}
		if newPassword != newPassword2 {
			fmt.Fprintln(w, "两次新密码不一致，请检查！")
			return
		}
		if password == newPassword {
			fmt.Fprintln(w, "新老密码一致，请检查！")
			return
		}
		if err := dao.ChangePWD(name, password, newPassword); err != nil {
			fmt.Fprintln(w, err.Error())
			return
		}
	}
}

func TokenLogin(w http.ResponseWriter, r *http.Request) {
	log.Log("method:", r.Method) //获取请求的方法
	if r.Method != "POST" {
		fmt.Fprintln(w, "只接受post请求")
		return
	}
	//t := template.New("test")
	//请求的是登陆数据，那么执行登陆的逻辑判断
	r.ParseForm()
	mapdata, err := util.GetBody(r)
	if err != nil{
		fmt.Fprintln(w, err.Error())
		return
	}

	userId, ok1 := mapdata["userId"].(string)
	userToken, ok2 := mapdata["userToken"].(string)
	ok := util.CheckOKs(ok1, ok2)
	if !ok {
		fmt.Fprintln(w, "params解析失败")
		return
	}
	if userId == "" || userToken == "" {
		fmt.Fprintln(w, "请填写账号、密码！")
		return
	}
	tokenErr := dao.CheckToken(userId, userToken)
	if tokenErr != nil {
		fmt.Fprintln(w, tokenErr)
		return
	}
	fmt.Fprintln(w, "token登录成功")
}

func SetDiscussReq(w http.ResponseWriter, r *http.Request){
	//defer r.Body.Close()
	//log.Log("SetDiscuss	! methos:	", r.Method)
	if r.Method != "POST" {
		fmt.Fprintln(w, "只接受post请求")
		log.Log("只接受Post请求！")
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
	fmt.Fprint(w, "操作成功1")
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

	log.Log("成功插入次数：	")
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
	log.Log(chatRoom, start, stop)
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
	for k, v := range res{
		log.Log(k, ": ", v)
	}
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