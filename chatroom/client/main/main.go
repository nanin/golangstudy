package main

import (
	"fmt"
	"os"
	"test/chatroom/client/processor"
)

var (
	userid   int
	userpwd  string
	username string
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
			fmt.Printf("请输入用户id号:")
			fmt.Scanf("%d\n", &userid)
			fmt.Printf("请输入用户密码:")
			fmt.Scanf("%s\n", &userpwd)
			err := processor.Login(userid, userpwd)
			if err != nil {
				fmt.Println(err)
				return
			}
			// err := login(userid, userpwd)
			// loop = false
		case 2:
			fmt.Printf("请输入用户id号:")
			fmt.Scanf("%d\n", &userid)
			fmt.Printf("请输入用户昵称:")
			fmt.Scanf("%s\n", &username)
			fmt.Printf("请输入用户密码:")
			fmt.Scanf("%s\n", &userpwd)
			err := processor.Register(userid, username, userpwd)
			if err != nil {
				fmt.Println(err)
				return
			}

		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("你的输入有误，重新输入！")

		}

		if !loop {
			break
		}
	}

}
