package processor

import (
	"encoding/json"
	"fmt"
	"net"
	"test/chatroom/common"
	"test/chatroom/server/model"
)

type Processor struct {
	Conn net.Conn
}

func (me *Processor) Process() {

	transfer := common.Transfer{
		Conn: me.Conn,
	}

	for {
		// fmt.Println("等待客户端连接。。。")
		msg, err := transfer.Get()
		if err != nil {
			// if err == io.EOF {
			// 	fmt.Printf("<End %s>因为客户端退出,msg=\n", me.Conn.LocalAddr())
			// }
			return
		}
		// fmt.Printf("<Begin %s>:%v\n", me.Conn.LocalAddr(), msg)
		switch msg.Type {
		case common.LoginMessageType:
			//如果是登录消息，抛给登录的Process处理
			var loginmsg common.LoginMessage
			err = json.Unmarshal([]byte(msg.Data), &loginmsg)
			if err != nil {
				return
			}
			userProcess := &UserProcess{
				Conn:        &me.Conn,
				ProcessUser: loginmsg,
			}
			userProcess.Login()
		case common.RegisterMessageType:
			var loginmsg common.LoginMessage
			err = json.Unmarshal([]byte(msg.Data), &loginmsg)
			if err != nil {
				return
			}
			userProcess := &UserProcess{
				Conn:        &me.Conn,
				ProcessUser: loginmsg,
			}
			userProcess.Register()
		case common.ChatRoomMessageType:
			var chatMsg common.ChatMessage
			_ = json.Unmarshal([]byte(msg.Data), &chatMsg)

			MessageProcess := &MessageProcess{
				Conn:     &me.Conn,
				FromUser: model.User{UserId: chatMsg.FromUserId},
			}
			if msg.Data != "" {
				MessageProcess.BroadCastMessage(chatMsg)
			}
			// fmt.Println(chatMsg)

		default:
			fmt.Printf("未检测到消息类型%v\n", msg.Type)
		}

	}
}
