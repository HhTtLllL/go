package main

import (
	"fmt"
	"sync"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {

	go func() {

		for i := 0; i < 10; i ++ {

			ch <- i
		}

		close(ch)				//关闭channel
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {

	go func () {

		for {

			if data, ok := <- ch; ok {

				fmt.Println(data)
			}else {

				//当ok为bool值的时候， 表示通道关闭，跳出循环
				break;
			}
		}
		wg.Done()
	}()

}
func main() {

	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	dataProducer(ch, &wg)

	wg.Add(1)
	dataReceiver(ch, &wg)

	wg.Add(1)
	dataReceiver(ch, &wg)

	wg.Wait()
}