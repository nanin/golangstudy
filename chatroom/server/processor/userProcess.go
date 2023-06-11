package processor

import (
	"encoding/json"
	"fmt"
	"net"
	"test/chatroom/common"
	"test/chatroom/server/model"
)

var OnlineUsersMap map[int]*UserProcess

type UserProcess struct {
	Conn        *net.Conn
	ProcessUser common.LoginMessage
}

func init() {
	OnlineUsersMap = make(map[int]*UserProcess, 1024)
}

func (me *UserProcess) Login() (err error) {

	var resultMsg common.Message
	resultMsg.Type = common.LoginResultMessageType
	var loginResul common.ResultMessage

	user, err := model.MyUserDao.GetByUserId(me.ProcessUser.UserId)
	if err != nil || user == nil {
		fmt.Println("MyUserDao.GetByUserId error:", err)
		loginResul.Code = 500
		loginResul.Error = "用户不存在"
		// return
	} else {

		if user.UserId == me.ProcessUser.UserId && user.UserPwd == me.ProcessUser.UserPwd {
			loginResul.Code = 200
			//加入到在线列表里去
			OnlineUsersMap[me.ProcessUser.UserId] = me

			fmt.Printf("%v登录成功,在线用户:", me.ProcessUser.UserId)
			for usr, _ := range OnlineUsersMap {
				fmt.Printf("%v,", usr)
			}
			fmt.Println()
		} else {
			loginResul.Code = 500
			loginResul.Error = "账号或者密码输入错误"
		}
	}
	// if loginmsg.UserId == 1 && loginmsg.UserPwd == "123456" {
	// 	loginResul.Code = 200
	// } else {
	// 	loginResul.Code = 500
	// 	loginResul.Error = "账号或者密码输入错误"
	// }

	//先把loginResult序列化,并赋值给resultMessage的data
	data, err := json.Marshal(loginResul)
	if err != nil {
		return

	}
	resultMsg.Data = string(data)
	transfer := common.Transfer{
		Conn: *me.Conn,
	}
	err = transfer.Send(resultMsg)
	return

}

func (me *UserProcess) Register() (err error) {

	var resultMsg common.Message
	resultMsg.Type = common.RegisterResultMessageType
	var resulMsg common.ResultMessage

	user := &model.User{
		UserId:   me.ProcessUser.UserId,
		UserName: me.ProcessUser.UserName,
		UserPwd:  me.ProcessUser.UserPwd,
	}

	err = model.MyUserDao.Insert(*user)
	if err != nil {
		resulMsg.Code = 500
		resulMsg.Error = err.Error()
	} else {
		resulMsg.Code = 200

	}

	data, err := json.Marshal(resulMsg)
	if err != nil {
		return

	}
	resultMsg.Data = string(data)
	transfer := common.Transfer{
		Conn: *me.Conn,
	}
	err = transfer.Send(resultMsg)
	return
}
