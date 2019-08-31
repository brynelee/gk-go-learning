package groutine_test

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine (t * testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()


	}

	time.Sleep(time.Millisecond * 50)
}
