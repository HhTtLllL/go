package main
// 13 min
import (
	"fmt"
	"os"
	"strings"

	//"strings"

	//"os"
	"net/http"
	"strconv"
	"regexp"
)

func HttpGet(url string)(result string, err error) {

	respon, err1 := http.Get(url)
	if err1 != nil {
		err = err1

		return
	}
	defer respon.Body.Close()

	//读取网页内容
	buf := make([]byte ,1024*4)

	for{
		n, _ := respon.Body.Read(buf)
		if n == 0{
			break;
		}

		result += string(buf[:n])
	}


	return
}


func SpiderPage(i int , page chan int){

	//https://www.pengfue.com/index_3.html
	url := "https://www.pengfue.com/index_" + strconv.Itoa(i) + ".html"
	fmt.Printf("开始爬取第%d页的网址%s\n",i,url)


	result, err := HttpGet(url)

	if err != nil{
		fmt.Println("Http get err = ",err)
		return
	}


	/*<h1 class="dp-b"><a href="https://www.pengfue.com/content_1857787_1.html" target="" +
		"_blank">系统维护的时候</a>*/

	//re := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*？))“`)
	re := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))" target="_blank">`)
	if re == nil{
		fmt.Println("regexp.MustCompile err")
		return
	}

	//fmt.Println("re = ", re)

	//fmt.Println("result = ", result)

	// 将字符串写入 文件中
	fileTitle := make([] string, 0)
	fileContent := make([] string, 0)


	//取出关键信息
	joyUrls := re.FindAllStringSubmatch(result, -1) //n 为正数， 查找前n个匹配项， n为 -1， 查找全部匹配项

	//fmt.Println("取出关键信息 = ",joyUrls)


	for _, data := range joyUrls {

		title, content, err := SpiderOneJoy(data[1])

		if err != nil{

			fmt.Println("SpiderOneJoy err = ", err)
			continue
		}

		fileTitle = append(fileTitle, title)
		fileContent = append(fileContent, content)
		//fmt.Println("title = #", title)
		//fmt.Println("content = #",content)
	}

	//把内容写入到文件中
	StoreJoyToFile(i, fileTitle, fileContent)

	page <- 1
	//fmt.Println(result)
}

func StoreJoyToFile(i int, title []string, content []string) {

	//新建文件
	filename, err := os.Create(strconv.Itoa(i) + ".txt")

	if err != nil{
		fmt.Println("os Create err = ", err)
		return
	}

	defer filename.Close()

	//写内容

	n := len(title)

	for i := 0 ; i < n ; i++ {

		//写标题
		filename.WriteString("《" + title[i] + "》" + "\n\n")
		//写内容
		filename.WriteString(content[i] + "\n")

		filename.WriteString( "\n=====================\n")


	}


}


func SpiderOneJoy(url string) (title, content string,err error ) {

	fmt.Println("进入SpiderOneJoy 函数")

	result, err1 := HttpGet(url)

	if err1 != nil {

		err = err1
		return
	}

	//取关键信息
	//取标题
	//<h1> 标题 /h1>

	re1 := regexp.MustCompile(`<h1>(?s:(.*?))/h1>`)
	if re1 == nil{
		err = fmt.Errorf("%s", "regexp.MustCompile err")
		return
	}

	//取内容
	tmpTitle := re1.FindAllStringSubmatch(result, 1)

	for _,data := range tmpTitle{

		title = data[1]

		title = strings.Replace(title, "\t", "", -1)
		title = strings.Replace(title, "<", "", -1)
		//fmt.Println("title == ",data[1])
		break
	}

	/*<div class="content-txt pt10">
		我：“人家都说打游戏是为了放松，你打游戏最放松的是什么时候？” 同事：“系统维护的时候。”<a id="prev"*/

	re1 = regexp.MustCompile(`<div class="content-txt pt10">(?s:(.*?))<a id="prev"`)

	if re1 == nil {
		err = fmt.Errorf("%s", "regexp.MustCompile err ")

		return
	}

	tmpmessage := re1.FindAllStringSubmatch(result, -1)

	for _, data := range tmpmessage{

		content = data[1]
		content = strings.Replace(content, "\t", "", -1)
		content = strings.Replace(content, "\n", "", -1)

		break
	}

	return
}


func DoWork(start, end int){

	page := make(chan int)
	fmt.Println("开始爬取第%d页到%d页的网址",start,end)
	for i := start; i <= end; i++ {
		go SpiderPage(i, page)
	}

	for i := start ; i <= end ; i++ {

		fmt.Println("第几页爬去完毕", <-page)
	}

}

func main(){

	var start, end int
	fmt.Printf("请输入爬取的起始页:")
	fmt.Scan(&start)

	fmt.Printf("请输入爬取的终止页:")
	fmt.Scan(&end)


	DoWork(start,end)

	return
}