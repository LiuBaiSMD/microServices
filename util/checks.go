package util

import (
	"fmt"
	"errors"
	"github.com/micro/go-micro/util/log"
	"os"
)

var allowedPassword string = "1234567890abcdefghijklmnopqrstuvwxyz!@#$%^&*()_+=-<>?,./:;'\"{}[]\\"
var allowedPasswordMap map[int32]int
var inited bool = false

func Init(){
	if inited == false{
		log.Log("初始化checks模块!")
		allowedPasswordMap = make(map[int32]int, 100)
		for _,v := range allowedPassword{
			allowedPasswordMap[v] = 1
		}
	}
	inited = true
}

func CheckPassword(password string) error {
	//检查密码的工具
	if (len(password) < 8 || len(password )> 20 ){
		fmt.Println("密码需在在8~20字符之间，请重设密码!")
		return errors.New("密码需在在8~20字符之间，请重设密码!")
	}
	if err := CheckIfInAllowed(password);err!= nil{
		return err
	}
	return nil
}

func CheckOKs(oks ...bool) bool{
	//检查oks是否全为true
	for _, v := range oks{
		if !v{
			return false
		}
	}
	return true
}

func CheckIfInAllowed(password string) error{
	for _, v := range password{
		if allowedPasswordMap[v] == 1{
			continue
		}else{
			notice := "出现未允许的字符！请使用:  " +  allowedPassword
			return errors.New(notice)
		}
	}
	return nil
}

func CheckReqAllowed(rMethod string, AllowedMethods ...string)bool{
	for _,v := range AllowedMethods{
		log.Log(v)
		if v == rMethod{
			return true
		}
	}

	return false
}

//检查路径是否存在
func CheckPathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//检查一个dir路径，没有则会创建
func CheckDirOrCreate(dirPath string) error{
	if ifExist,err := CheckPathExists(dirPath); err != nil{
		return err
	}else if !ifExist{
		err1 := os.MkdirAll(dirPath, 0777)
		if err1!=nil{
			return err1
		}
	}
	return nil
}