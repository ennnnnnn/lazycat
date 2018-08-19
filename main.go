package main

import (
	."./framework/client"
	"./framework/service"
	"net/http"
	"fmt"
)

func main() {
	service.HandleStatic("/favicon.ico", "")
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		c := GetClient(req)
		fmt.Println("IP地址：" + c.IPaddress())
		gv := c.GetValue()
		fmt.Println(gv.Get("test").Int())
		fmt.Println("电子邮件：" + gv.Get("email").Email())
	})
	fmt.Println("run.....")
	http.ListenAndServe(":80", nil)
}