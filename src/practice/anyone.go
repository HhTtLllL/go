package main

import (
	"fmt"
	"runtime"
	"time"
)

func runTask(id int) string {

	time.Sleep(10 * time.Millisecond)

	return fmt.Sprintf("The result is from %d", id)
}

func FirstResponse() string {

	numOfRunner := 10
	ch := make(chan string, numOfRunner)

	for i := 0; i < numOfRunner; i ++ {

		go func(i int) {

			ret := runTask(i)
			ch <- ret
		}(i)
	}

	return <- ch
}

func main() {

	fmt.Println("before:", runtime.NumGoroutine())
	fmt.Println(FirstResponse())
	time.Sleep(time.Second * 1)
	fmt.Println("before:", runtime.NumGoroutine())

}