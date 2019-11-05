package crontabTask

import (
	"time"
	"testing"
	"fmt"
)
func TestAddTask(t *testing.T) {
	go TimeRun("taskNow", Task1,2,"1", "a", []int{1,2,3})
	var over chan int
	over = make(chan int)
	go func(){
		tickOver := time.NewTicker(time.Second * 13)
		tick := time.NewTicker(time.Second )
		for {
			select{
			case <- tick.C:
				RunTaskNow("taskNow")
			case <- tickOver.C:
				over <- 1
			}
		}
	}()
	<-over
	close(over)
	fmt.Println("over")
}

func Task1(params ...interface{})([]interface{}, error){
	fmt.Println("Task1", params)
	return nil, nil
}