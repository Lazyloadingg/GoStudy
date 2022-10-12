package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

const port = ":9090"

func main() {

	r := gin.Default()
	//gin文件托管
	r.StaticFS("files", http.Dir("./files"))
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", uploadPage)
	r.POST("/upload", upload)
	r.Run(port)
}

// 上传页
func uploadPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

// 上传
func upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	path := ctx.Request.URL.RawPath
	if err != nil {
		ctx.JSON(http.StatusOK, err)
	} else {
		dst := "./files/" + file.Filename
		err = ctx.SaveUploadedFile(file, dst)
		if err != nil {
			ctx.JSON(http.StatusOK, err)
		} else {

			fmt.Printf("ctx.Request.URL.RawPath: %v\n", path)
			url := GetOutboundIP() + port + "/files/" + file.Filename
			ctx.JSON(http.StatusOK, gin.H{
				"code": 200,
				"url":  url,
				"msg":  "success",
			})
		}
	}
}

// 获取本机ip
func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.String())
	return localAddr.IP.String()
}

func getIp() string {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())
	}
	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return ipnet.String()
					}
				}
			}
		}
	}
	return ""
}
