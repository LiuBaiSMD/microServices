package crontabTask

import (
	"time"
	"testing"
	"fmt"
)
func TestAddTask(t *testing.T) {

	//生产端
	go TimeRun("taskNow", Task1,2,"1", "a", []int{1,2,3})

	//消费端
	go func(){
		endPoint:
		tickOver := time.NewTicker(time.Second * 10)
		tick := time.NewTicker(time.Second )
	for {
		select{
		case <- tick.C:
			RunTaskNow("taskNow")
		case <- tickOver.C:
			goto endPoint
		}
	}
	}()

	time.Sleep(time.Second * 11)
}

func Task1(params ...interface{})([]interface{}, error){
	fmt.Println("Task1", params)
	return nil, nil
}