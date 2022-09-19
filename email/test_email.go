package main

import (
	"crypto/tls"
	"fmt"

	"gopkg.in/gomail.v2"
)

func main() {
	sendEmail()
}

// 发送邮件以我个人常用的163邮箱为例
func sendEmail() {

	host := "smtp.163.com"
	port := 25
	userName := "xxx@163.com" //用户名
	passWord := ""            //密码
	fmt.Printf("\"发送邮件\": %v\n", "发送邮件")

	content := "邮件内容"

	m := gomail.NewMessage()         //生成信息
	m.SetHeader("From", userName)    //发件人
	m.SetHeader("To", "xxx@163.com") //收件人
	m.SetHeader("Subject", "测试邮件")   //主题
	m.SetBody("text/plain", content) //内容

	d := gomail.NewDialer(
		host, port, userName, passWord,
	) //创建发送器

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true} //关闭tls验证
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}
