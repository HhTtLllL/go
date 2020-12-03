package main

import (
	"fmt"
	"time"
)

func service() string {

	time.Sleep(time.Millisecond * 50)

	return "Done"
}
func otherTask() {

	fmt.Println("work on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is Done")
}

func AsynService() chan string {

	retCh := make(chan string)
	go func() {

		//time.Sleep(time.Millisecond * 50)
		ret := service()
		fmt.Println("returned result.")
		time.Sleep(time.Second * 2)
		retCh<- ret
		fmt.Println("service exited.")
	}()

	fmt.Println("Asynservice .....")
	return retCh
}

func main() {

	retCh := AsynService()
	otherTask()
	fmt.Println(<-retCh)
}