package select_test

import (
	"time"
	"fmt"
	"testing"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "service: Done"
}

func AsyncService() chan string {
	retCh := make(chan string)
	go func() {
		ret := service()
		fmt.Println("AsyncService child thread: returned result")
		//time.Sleep(time.Millisecond * 50)
		retCh <- ret
		fmt.Println("AsyncService internal thread: service exited.")
	}()
	return retCh
}

func TestSelect(t *testing.T){
	select {
		case ret := <-AsyncService():
			t.Log(ret)
		case <-time.After(time.Millisecond *  100):
			t.Error("time out")
	}
}