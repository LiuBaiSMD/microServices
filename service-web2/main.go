package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	// 设置路由，如果访问/，则调用index方法
	http.Handle("/css/", http.FileServer(http.Dir("html")))
	http.Handle("/js/", http.FileServer(http.Dir("html")))
	http.HandleFunc("/", love)

	// 启动web服务，监听9090端口
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
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

