package dao

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/util/log"
	"myOwnStation/util"
	"time"
)

var sqlDB *gorm.DB

type Auth struct {
	Id       int 	`gorm:"AUTO_INCREMENT;primary_key;"`
	UserId	 string `gorm:"column:userid;"`
	Password string	`gorm:"column:password"`
	Name	 string
	PhoneNum int	`gorm:"column:phoneNum"`
	Age 	 int
	CreateAt	string `gorm:"column:createAt`
	Num          int     `gorm:"AUTO_INCREMENT"`
}

func InitMysql(driverName, mysqlURL string) (*gorm.DB) {
	if !inited{
		util.Init()
		var err error
		sqlDB, err = gorm.Open(driverName, mysqlURL)
		if err != nil {
			log.Log(err)
			panic("failed to connect database")
		}
		createMyTable()
	}
	//GetCreateAt()
	inited = true
	return sqlDB
}

func createMyTable() {
	//初始化Auth,并且插入一个成员
	sqlDB.DropTable(&Auth{})
	if !sqlDB.HasTable(&Auth{}){
		//创建并更新表结构
		log.Log("创建并更新表结构")
		sqlDB.Set("gorm:table_options", "ENGINE=InnoDB AUTO_INCREMENT=1;").CreateTable(&Auth{})
	}
	sqlDB.AutoMigrate(&Auth{})
}

func QueryUserIdPass(userId, password string) error {
	userAuth := Auth{
		UserId: userId,
	}
	res := Auth{}
	if sqlDB.Find(&Auth{}, userAuth).Scan(&res).RecordNotFound(){
		return errors.New("无此用户，请进行注册")
	}else if res.Password != password {
		return errors.New("密码错误，请重试")
	}
	return nil
}

func RegisterUserIdPass(userId, password string) error {
	//0 操作成功  1 密码不合格 2 账号已被注册
	//var users []Auth
	userAuth := Auth{
		UserId: userId,
	}
	res := Auth{}
	log.Log(userId, "	:	", password)
	log.Log("userAuth:	", userAuth.UserId, userAuth)
	if !sqlDB.Find(&Auth{}, userAuth).Scan(&res).RecordNotFound(){
		log.Log("res:	", res)
		return errors.New("此账户已注册，请重新输入")
	}
	if err := util.CheckPassword(password); err != nil {
		return err
	}
	//if res.UserId != "" {
	//	return errors.New("此账户已注册，请重新输入")
	//}
	userAuth.Password = password
	userAuth.Name = "test"
	userAuth.CreateAt = time.Now().String()
	//sqlDB.Create(&userAuth)
	if !sqlDB.NewRecord(userAuth){
		return errors.New("创建失败！")
	}
	log.Log("创建成功！")
	return nil
}

func ChangePWD(userId, password, newPassword string) error {
	if password == newPassword {
		return errors.New("新老密码一致，请重新填写！")
	}
	//检查用户是否存在，不存在不让改密码
	userAuth := Auth{
		UserId: userId,
	}
	res := Auth{}
	if sqlDB.Find(&Auth{}, userAuth).Select("userid").Scan(&res).RecordNotFound(){
		return errors.New("此账户未注册，请重新输入")
	}
	if err := util.CheckPassword(password); err != nil {
		return err
	}
	if err := QueryUserIdPass(userId, password); err != nil {
		return err
	}
	res.Password = newPassword
	sqlDB.Model(&res).Update("password", newPassword)
	return nil
}

func GetCreateAt(){
	userAuth := Auth{
		UserId: string(1),
	}
	sqlDB.Find(&Auth{}, userAuth).Scan(&userAuth)
	//times := userAuth.CreateAt
	//log.Log(times.Second())
}
