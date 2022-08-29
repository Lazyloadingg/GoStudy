package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	httpClient()
}
func httpClient() {

	resp, err := http.Get("https://www.baidu.com")
	body, err := ioutil.ReadAll(resp.Body) //ioutil提供的方法标准库可以读取整个文件内容
	if err != nil {
		fmt.Printf("read from resp.Body failed, err:%v\n", err)
		return
	}
	fmt.Print(string(body))

	defer resp.Body.Close()

}
