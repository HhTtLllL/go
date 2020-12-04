package main

import (
	"fmt"
	"sync"
	"unsafe"
)

var singleInstance *Singleton
var once sync.Once

type Singleton struct {

}


func GetSinglketonObj() *Singleton {

	once.Do(func () {

		fmt.Println("creat Obj")
		singleInstance = new(Singleton)
	})

	return singleInstance
}


func main() {

	var wg sync.WaitGroup

	for i := 0; i < 10; i ++ {

		wg.Add(1)
		go func () {

			obj := GetSinglketonObj()
			fmt.Printf("%x\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}

	wg.Wait()
}