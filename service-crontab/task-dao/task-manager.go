package taskdao

import (
	"github.com/go-redis/redis"
	//"time"
)

type TaskManeger struct {
	client *redis.Client
}

//type TaskOptions interface {
//	AddTask(taskStr string)error
//	RmTask(taskStr string)error
//	GetTask(timeStamps string)error
//}
//
//func (t *TaskManeger) AddTask(taskStr string)error{
//	client.LPush(TaskKey, taskStr)
//}