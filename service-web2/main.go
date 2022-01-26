}

func eatNotify(){
        for{
                if time.Now().Hour()==17 && time.Now().Minute()==42{
                        s := randowChoice(notifyEatList)
                        sendNotifyToMyWife(testBarkUrl, s)
                        }
                fmt.Println(time.Now().Hour(), time.Now().Minute())
                time.Sleep(60*time.Second)
        }

}

func notify(h, m int, l []string){
        for{
                if time.Now().Hour()==(h-8) && time.Now().Minute()==m{
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
