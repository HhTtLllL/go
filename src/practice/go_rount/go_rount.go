package main

import (
	"fmt"
	"time"
)

/* 绝大多数情况下是 无输出. 因为主协程先结束, 整个程序就结束

	第二种情况:10 个 10, 当 for 循环这个语句执行完的时候, 一个go 协程都没有开始, 此时i的值已经变为 10,在执行 go grount 会输出10个10

	第三种情况:0-9乱序输出0-9
	等等之类的结果



*/
func main() {

	for i := 0; i < 10; i ++ {

		go func(i int) {

			fmt.Println(i)
		}(i)
	}

	time.Sleep(time.Millisecond * 500)
}