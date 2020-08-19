package main

import (
	"fmt"
	"strconv"
	"net/http"
	"os"
)

func HttpGet(url string) (result string, err error){

	resp, err1 := http.Get(url)
	if err1 != nil {

		err = err1

		return
	}

	defer resp.Body.Close()

	buf := make([]byte, 1024*4)

	for{
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("read err = ",err)
			break
		}
		result += string(buf[:n])
	}

	return

}

func DoWork(start, end int){

	for i:= start; i <= end; i++ {

		url := "https://tieba.baidu.com/f?kw=lol&ie=utf-8&pn"+strconv.Itoa((i - 1)*50)

		fmt.Println("url = ", url)

		result, err := HttpGet(url)

		if(err != nil){
			fmt.Println("HttpGet err = ",err)
			continue
		}

		fileName := strconv.Itoa(i) + ".html"
		f, err1 := os.Create(fileName)
		if err1 != nil {
			fmt.Println("os.Creae err1 = ", err1)
			continue
		}

		f.WriteString(result)

		f.Close()

	}

}

func main(){

	var start, end int

	fmt.Scan(&start)
	fmt.Scan(&end)

	DoWork(start, end)
}
