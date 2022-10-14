package main

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

const port = ":9090"

type Manifest struct {
	XMLName xml.Name `xml:plist`
	Version string   `xml:"version,attr"`
	Dict    string   `xml:",innerxml"`
}

func main() {

	// r := gin.Default()
	// //gin文件托管
	// r.StaticFS("files", http.Dir("./files"))
	// r.LoadHTMLFiles("./index.html")
	// r.GET("/index", uploadPage)
	// r.POST("/upload", upload)
	// r.Run(port)

	// file, err := ioutil.ReadFile("./manifest.plist")
	// if err != nil {
	// 	fmt.Printf("\"plist读取出错\": %v\n", "plist读取出错")
	// 	return
	// }

	// var model Manifest
	// err = xml.Unmarshal(file, &model)
	// if err != nil {
	// 	fmt.Printf("\"xml序列化失败\": %v\n", "xml序列化失败")
	// 	return
	// }

	// fmt.Printf("model: %v\n", model)

	// doc := etree.NewDocument()
	// if err := doc.ReadFromFile("./Payload/CollegeEnglishTest.app/info.plist"); err != nil {
	// 	fmt.Printf("\"etree读取失败\": %v\n", "etree读取失败")
	// 	return
	// }

	// ReadInfoPList("./Payload/CollegeEnglishTest.app/info.plist")
	CreatePlist()

}
func findDir(dir string) string {
	fileinfo, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	// 遍历这个文件夹
	for _, fi := range fileinfo {

		// 重复输出制表符，模拟层级结构

		// 判断是不是目录
		if fi.IsDir() {
			println(`目录：`, fi.Name())
			findDir(dir + `/` + fi.Name())
			ok := strings.HasSuffix(fi.Name(), ".plist")
			if ok {
				fmt.Printf("\"找到了\": %v\n", fi.Name())
				return fi.Name()
			}
		} else {
			return ""
		}
	}
	return ""
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

// Unzip decompresses a zip file to specified directory.
// Note that the destination directory don't need to specify the trailing path separator.
func Unzip(zipPath, dstDir string) error {
	// open zip file
	reader, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		if err := unzipFile(file, dstDir); err != nil {
			return err
		}
	}
	return nil
}

func unzipFile(file *zip.File, dstDir string) error {
	// create the directory of file
	filePath := path.Join(dstDir, file.Name)
	if file.FileInfo().IsDir() {
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			return err
		}
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	// open the file
	rc, err := file.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	// create the file
	w, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer w.Close()

	// save the decompressed file content
	_, err = io.Copy(w, rc)
	return err
}
