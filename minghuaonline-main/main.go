package main

import (
	"fmt"
	"minghuaonline/api"
	"os"
	"time"
)

func main() { //主函数
	var an string
ou:
	fmt.Printf("请输入开启密码：")
	// fmt.Scanf("%s\n", &an)
	timeStr := time.Now().Format("2006-01-02 15:04:05") //当前时间的字符串，2006-01-02 15:04:05据说是golang的诞生时间，固定写法
	times := "2023-12-05 17:43:00"
	if timeStr >= times {
		fmt.Println("超时", timeStr)
		os.Exit(0)
	}
	fmt.Println("当前时间", timeStr) //打印结果：2017-04-11 13:24:04
	fmt.Println("过期时间", times)   //打印结果：2017-04-11 13:24:04

	an = "666"
	switch an {
	case "666":
		fmt.Printf("密码输入正确,请开始刷课\n")
		api.GetSchool()
		api.Login()
	default:
		fmt.Printf("密码输入错误！\n")
		goto ou
	}
}
