package concurrency_test

import (
	"time"
	"fmt"
	"testing"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "service: Done"
}

func otherTask() {
	fmt.Println("otherTask: working on something else...")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("otherTask: Task is done.")
}

func TestService(t *testing.T){
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("AsyncService child thread: returned result")
		retCh <- ret
		fmt.Println("AsyncService internal thread: service exited.")
	}()
	return retCh
}


func TestAsyncService(t *testing.T){

	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
	time.Sleep(time.Second * 1)
}

