package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	//原生net框架实现web服务器
	// http.HandleFunc("/", handle)
	// err := http.ListenAndServe(":9000", nil)
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// }

	// res, err := http.Get("https://www.baidu.com")
	// if err == nil {
	// 	fmt.Printf("百度res: %v\n", res)
	// }

	r := gin.Default()
	r.GET("/", ginGet)
	r.Run(":9000")

}

func ginGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "白日依山尽",
	})
}

// func handle(res http.ResponseWriter, req *http.Request) {
// 	req.ParseForm()
// 	// fmt.Printf("req: %v\n", req)
// 	fmt.Printf("req.URL: %v\n", req.URL)
// 	fmt.Printf("req.Host: %v\n", req.Host)
// 	fmt.Printf("req.Body: %v\n", req.Body)
// 	fmt.Printf("req.Method: %v\n", req.Method)
// 	res.Write([]byte("<h1>hello world</h1>"))
// }
