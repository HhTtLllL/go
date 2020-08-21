package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)



func main(){

	url := "https://blog.csdn.net/qq_43701555"
	url_list := url + "/article/list/"

	fmt.Println("开始爬取")

	count := 0

	i := 1
	for {

		url_tmp := url_list
		url_tmp += strconv.Itoa(i)

		fmt.Println("i = ",i)

		fmt.Println("url = ", url_tmp)

		n := dowok(url_tmp)

		if  n == 0 {
			fmt.Println("全部爬去完毕")
			break
		}

		count += n

		i++
	}

	fmt.Println("共爬取 ", count, "篇博客")

}

func dowok(url_list string) int{

	n := 0

	result := httpGet(url_list)

	if strings.Contains(result, "空空如也"){

		fmt.Println("空空如也")
		return n
	}

	ever_url := deal(result)

	ever_only_url  := deal_2(ever_url)


	/*for _, data := range ever_url{

		fmt.Printf("data = %s\n",data)
		//return 0
	}*/

	count := 0;

	for i := 0; i < len(ever_only_url); i++ {

		if(ever_only_url[i] == "" ) {

			break
		}
		count++
		fmt.Println("len1 = ", len(ever_only_url[i]))
		//fmt.Println(ever_only_url[i])
	}

	//fmt.Println(i , "\n")
	return count
}

func httpGet(url_list string) (result string){

	respon, err := http.Get(url_list)

	if err != nil {

		fmt.Println("http.get err = ", err)

		return
	}

	defer respon.Body.Close()

	//读取网页内容

	buf := make([] byte, 4*1024)

	for {

		n, _ := respon.Body.Read(buf)

		if n == 0 {
			break
		}

		result += string(buf[:n])
	}

	return
}

//第一次筛选
func deal(result string) ([][] string){

	//<a href="https://blog.csdn.net/qq_43701555/article/details/102920880" target="_blank">

	var url [][]string

	re := regexp.MustCompile(`<p class="content">(?s:(.*?))" target="_blank">`)

	//re := regexp.MustCompile(`<a href="(?s:(.*?))" target="_blank">`)
//	fmt.Printf("re = %s\n",re)

	if re == nil {
		fmt.Println("没有在主页中找到博客链接")
		return url
	}


	url = re.FindAllStringSubmatch(result, -1)

	return url
}


//第二次 筛选
func deal_2(ever_url [][]string) []string {

	fmt.Println("len = ",len(ever_url))
	ever_only_url := make([]string, 40)

	for i, data := range ever_url{

		//fmt.Printf("data = %s\n",data[1])
		ever_only_url[i] = data[1][16:]
		//ever_only_url[i] = data[1]
	}

	//fmt.Println(ever_only_url[0])

	return ever_only_url
}
