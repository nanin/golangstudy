package processor

import (
	"encoding/json"
	"fmt"
	"net"
	"test/chatroom/common"
)

func backgroundProcess(con net.Conn) {
	tf := &common.Transfer{
		Conn: con,
	}
	for {
		// fmt.Println("客户端正在读取服务器消息")
		// fmt.Printf("<我>：")
		msg, err := tf.Get()
		if err != nil {
			fmt.Println("tf.Get 异常：", err)
			return
		}

		switch msg.Type {
		case common.ChatRoomMessageType:
			var chatMsg common.ChatMessage
			_ = json.Unmarshal([]byte(msg.Data), &chatMsg)
			fmt.Printf("<%v>：%v\r\n", chatMsg.FromUserId, chatMsg.Data)

		default:
			fmt.Printf("未检测到消息类型%v\n", msg.Type)
		}

		//读取到就下一步逻辑
		// fmt.Printf("读取到的消息体：%v", msg)
	}
}
