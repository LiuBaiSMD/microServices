package access

import (
	"fmt"
	r "github.com/go-redis/redis"
	"github.com/LiuBaiSMD/microServices/base/dao"
	"sync"
)

var (
	s  *service
	ca *r.Client
	m  sync.RWMutex
)

// service 服务
type service struct {
}

// Service 用户服务类
type Service interface {
	// MakeAccessToken 生成token
	MakeAccessToken(subject *Subject) (ret string, err error)

	// GetCachedAccessToken 获取缓存的token
	GetUserAccessToken(subject *Subject) (ret string, err error)

	// DelUserAccessToken 清除用户token
	DelUserAccessToken(token string) (err error)
}

// GetService 获取服务类
func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return s, nil
}

// Init 初始化用户服务层
func Init() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}
	dao.Init()
	ca = dao.GetRedis()

	s = &service{}
}
