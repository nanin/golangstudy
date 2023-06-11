package processor

import (
	"encoding/json"
	"fmt"
	"net"
	"test/chatroom/common"
)

func Login(userid int, userPwd string) (err error) {

	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return
	}
	defer conn.Close()

	transfer := common.Transfer{
		Conn: conn,
	}

	loginMsg := common.LoginMessage{
		UserId:  userid,
		UserPwd: userPwd,
	}
	data, err := json.Marshal(loginMsg)
	if err != nil {
		return
	}

	msg := common.Message{
		Type: common.LoginMessageType,
		Data: string(data),
	}
	err = transfer.Send(msg)
	if err != nil {
		return
	}
	msg, err = transfer.Get()
	if err != nil {
		return
	}
	if msg.Type == common.LoginResultMessageType {
		var loginResul common.ResultMessage
		json.Unmarshal([]byte(msg.Data), &loginResul)
		if loginResul.Code == 200 {
			fmt.Println("======登录成功，成功进入聊天室！！！！======")
			fmt.Println("<系统>可输入\"/?\" 获取帮助")

			//创建一个携程和服务器保持沟通
			go backgroundProcess(conn)

			MsgProcess := &MessageProcess{
				Conn:          conn,
				CurrentUserId: userid,
			}

			MsgProcess.ListenInput()
			//登录成功后，显示登录成功后的菜单
			// for {
			// 	showMenu()
			// }

		} else {
			fmt.Printf("账号'%v'登录失败，%s\n", userid, loginResul.Error)
		}
	} else {
		fmt.Printf("返回类型错误%v\n", msg.Type)
	}

	return

}

func Register(userid int, username string, userPwd string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return
	}
	defer conn.Close()

	transfer := common.Transfer{
		Conn: conn,
	}

	loginMsg := common.LoginMessage{
		UserId:   userid,
		UserPwd:  userPwd,
		UserName: username,
	}
	data, err := json.Marshal(loginMsg)
	if err != nil {
		return
	}
	msg := common.Message{
		Type: common.RegisterMessageType,
		Data: string(data),
	}
	err = transfer.Send(msg)
	if err != nil {
		return
	}
	msg, err = transfer.Get()
	if err != nil {
		return
	}
	if msg.Type == common.RegisterResultMessageType {
		var loginResul common.ResultMessage
		json.Unmarshal([]byte(msg.Data), &loginResul)
		if loginResul.Code == 200 {
			fmt.Println("注册成功")

		} else {
			fmt.Printf("账号'%v'注册失败，%s\n", userid, loginResul.Error)
		}
	} else {
		fmt.Printf("返回类型错误%v\n", msg.Type)
	}

	return
}

// func showMenu() {
// 	fmt.Println("是否进入聊天室(Y/N)")
// 	var key string
// 	fmt.Scanln(&key)
// 	switch key {
// 	case "Y", "y":
// 		fmt.Println("进入聊天室。。。。。。。")
// 	}
// }
