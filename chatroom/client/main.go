package main

import (
	"fmt"
	"os"
)

var (
	userid  int
	userpwd string
)

func main() {

	//接受用户选择
	var key int
	//判断是否接受显示菜单
	var loop = true

	for {
		fmt.Println("------------欢迎登录多人聊天系统------------")
		fmt.Println("\t\t 1.登录聊天室")
		fmt.Println("\t\t 2.注册用  户")
		fmt.Println("\t\t 3.退出系统")
		fmt.Println("请选择(1~3):")

		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("...正在登录聊天室...")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
		default:
			fmt.Println("你的输入有误，重新输入！")
			os.Exit(0)
		}

		if !loop {
			break
		}
	}

	//根据用户输入，显示新信息
	if key == 1 {
		//说明用户需要登录了
		fmt.Printf("请输入用户id号:")
		fmt.Scanf("%d\n", &userid)
		fmt.Printf("请输入用户密码:")
		fmt.Scanf("%s\n", &userpwd)
		err := login(userid, userpwd)
		if err != nil {
			fmt.Println("登录失败")
		} else {
			fmt.Println("登录成功")
		}
	} else if key == 2 {
		//用户注册逻辑
	}

}
