package main

/**
 *写爬虫
 */
import (
	"log"
	"regexp"
	"github.com/Unknwon/com"
	"net/http"
	"strings"
	"path"
)

var  imgs  = regexp.MustCompile(`src=\*http(.*?).jpg`)
func main() {

	num := make(chan bool,10)
	for i :=0; i<10; i++ {

		data, err := com.HttpGetBytes(&http.Client{}, "http://www.youku.com", nil)
		if err != nil {
			log.Fatal("获取页面失败 {%d},%v", 0, err)
		}

		matches := imgs.FindAll(data, -1)
		for _, match := range matches {
			num<-true
			log.Println(string(match))
			go download(string(match),num)
		}
	}

}

func download(url string,num chan bool)  {
	url = strings.TrimPrefix(url,"src=")
	log.Printf("正在下载 %s",url)
	err := com.HttpGetToFile(&http.Client{},url,nil,"pics/"+path.Base(url))
	if err !=nil {
		log.Printf("下载失败 %s,%s",url,err)
	}
	<-num
}
