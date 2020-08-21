package main

import (
	"net/http"
	"os"
)

func HttpGet(url string) (result string){

	resp,_ := http.Get(url)

	defer resp.Body.Close()

	buf := make([]byte, 4*1024)

	for{
		n,_ := resp.Body.Read(buf)

		if n == 0 {
			break
		}

		result += string(buf[:n])
	}

	return
}

func doWork(){

	url := "https://blog.csdn.net/qq_43701555/"
	result := HttpGet(url)

	fileName := "HhTtLllL_csnd" + ".html"

	f,_ := os.Create(fileName)

	f.WriteString(result)

	f.Close()




}
func main(){

	doWork()
}
