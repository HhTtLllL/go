package main

import (
	"fmt"
	"runtime"
	"time"
)

func runTask1(id int) string {

	time.Sleep(10 * time.Millisecond)

	return fmt.Sprintf("The result is from %d", id)
}

func AllFirstResponse() string {

	numOfRunner := 10
	ch := make(chan string, numOfRunner)

	for i := 0; i < numOfRunner; i ++ {

		go func(i int) {

			ret := runTask1(i)
			ch <- ret
		}(i)
	}

	finalRet := ""
	for j := 0; j < numOfRunner; j ++ {

		finalRet += <- ch + "\n"
	}

	return finalRet
}

func main() {

	fmt.Println("before:", runtime.NumGoroutine())
	fmt.Println(AllFirstResponse())
	time.Sleep(time.Second * 1)
	fmt.Println("before:", runtime.NumGoroutine())

}