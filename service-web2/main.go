package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var barkUrl = "https://api.day.app/V8bF9xQkVFTeV5otqvCEbT/"
var testBarkUrl = "https://api.day.app/rpiG6xZWiFxYHh9jg9fdZV/"

var notifyEatList = []string{"老婆老婆，吃饭了~","老婆，不要饿到肚子，吃饭了~", "老婆，快吃饭啦~"}
var notifyMorningList = []string{"早安老婆~","老婆，起床时间到了，早餐要记得次~", "老婆，吃饭饭了~"}
var notifyNightList = []string{"晚安老婆~","老婆，起准备睡觉啦~", "老婆，不要熬到太晚~"}

func main() {

	go notify(17, 35, notifyEatList)
	go notify(8, 8, notifyMorningList)
	go notify(23, 58, notifyNightList)
	// 设置路由，如果访问/，则调用index方法
	http.Handle("/css/", http.FileServer(http.Dir("html")))
	http.Handle("/js/", http.FileServer(http.Dir("html")))
	http.Handle("/flac/", http.FileServer(http.Dir("html")))
	http.HandleFunc("/", love)

	// 启动web服务，监听9090端口
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func eatNotify(){
	for{
		if time.Now().Hour()==17 && time.Now().Minute()==42{
			s := randowChoice(notifyEatList)
			sendNotifyToMyWife(barkUrl, s)
			}
		fmt.Println(time.Now().Hour(), time.Now().Minute())
		time.Sleep(60*time.Second)
	}

}

func notify(h, m int, l []string){
	for{
		if time.Now().Hour()==h && time.Now().Minute()==m{
			s := randowChoice(l)
			sendNotifyToMyWife(barkUrl, s)
		}
		fmt.Println(time.Now().Hour(), time.Now().Minute())
		time.Sleep(60*time.Second)
	}

}

func love(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	log.Println("method:", r.Method) //获取请求的方法
	log.Println("加载登录界面")
	t, err := template.ParseFiles("html/main.html")
	log.Println("err: ", err)
	t.Execute(w, nil)
	log.Println("加载完毕")
	return

}

func sendNotifyToMyWife(barkUrl, msg string) {
	//定时推送提醒给我老婆
	url := barkUrl + msg
	httpClient := &http.Client{Timeout:5 *time.Second}
	fmt.Println(url)
	httpClient.Get(url)
}

func Get(url string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}


func randowChoice(l []string) string {
	//随机从中间选一个
	n := len(l)
	s := rand.Intn(n)
	ss := l[s]
	return ss
}
