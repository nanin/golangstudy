package processor

import (
	"encoding/json"
	"fmt"
	"net"
	"test/chatroom/common"
)

type MessageProcess struct {
	Conn          net.Conn
	CurrentUserId int
}

func (me *MessageProcess) ListenInput() {
	var input string
	for {
		// fmt.Printf("<我>：")
		fmt.Scanln(&input)
		me.sendMsg(input)
	}
}

func (me *MessageProcess) sendMsg(input string) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return
	}
	defer conn.Close()

	transfer := common.Transfer{
		Conn: conn,
	}

	chatMsg := common.ChatMessage{
		FromUserId: me.CurrentUserId,
		Data:       input,
	}
	data, _ := json.Marshal(chatMsg)

	msg := common.Message{
		Type: common.ChatRoomMessageType,
		Data: string(data),
	}
	err = transfer.Send(msg)
	if err != nil {
		return
	}
}
