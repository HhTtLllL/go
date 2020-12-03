package main

import (
	"fmt"
	"time"
)

func isCancelled(cancelCh chan struct{}) bool {

	select {

	case <- cancelCh:
		return true
	default:
		return false
	}
}

func cancel_1(cancelCh chan struct {}) {

	cancelCh <- struct{}{}
}

func cancel_2(cancelCh chan struct{}) {

	close(cancelCh)
}


func main() {

	cancelChan := make(chan struct{}, 0)

	for i := 0; i < 5; i ++ {

		go func(i int, cancelCh chan struct{}) {

			for {

				//在循环里面不断判断任务是否被取消
				if isCancelled(cancelCh) {

					break
				}
				time.Sleep(time.Millisecond * 50)
			}
			fmt.Println(i, "Done")
		}(i, cancelChan)
	}

	cancel_2(cancelChan)
	time.Sleep(time.Second * 1)
}