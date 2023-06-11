package processor

import (
	"encoding/json"
	"fmt"
	"net"
	"test/chatroom/common"
	"test/chatroom/server/model"
)

type MessageProcess struct {
	Conn     *net.Conn
	FromUser model.User
}

func (me *MessageProcess) BroadCastMessage(msg common.ChatMessage) {

	//1、先拿到在线用户
	//2、for循环在线用户，进行发送消息（需要去除自己）

	for user, val := range OnlineUsersMap {
		if user == me.FromUser.UserId {
			continue
		}

		transfer := common.Transfer{
			Conn: *val.Conn,
		}

		data, _ := json.Marshal(msg)
		transfer.Send(common.Message{
			Type: common.ChatRoomMessageType,
			Data: string(data),
		})

		fmt.Println(string(data))
	}
}
