package common

import (
	"encoding/binary"
	"encoding/json"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [6089]byte
}

/*
1. 先发个head，确认有几个字节
2. 发送真正的消息体

*/
func (me *Transfer) Send(msg Message) (err error) {

	//把传入的Message进行序列化
	data, err := json.Marshal(msg)
	if err != nil {
		return
	}
	msgLen := len(data)
	binary.BigEndian.PutUint32(me.Buf[:4], uint32(msgLen))

	//先发送一个长度（Message的长度）给对方
	_, err = me.Conn.Write(me.Buf[:4])
	if err != nil {
		return
	}

	// fmt.Printf("data %s", string(data))
	//发送消息体
	// me.Buf = data[:msgLen]
	n, err := me.Conn.Write(data[:msgLen])
	if n != int(msgLen) || err != nil {
		return
	}
	return
}

func (me *Transfer) Get() (msg Message, err error) {
	// fmt.Println("正在读取数据")
	//先读到长度 放到msgLen里去
	data := make([]byte, 4)
	_, err = me.Conn.Read(data)
	if err != nil {
		return
	}
	msgLen := binary.BigEndian.Uint32(data)

	//根据msgLen读取消息体的真正内容,放到Buf里去
	n, err := me.Conn.Read(me.Buf[:msgLen])
	if n != int(msgLen) || err != nil {
		return
	}

	//把buff里的真正内容反序列化成Message
	err = json.Unmarshal(me.Buf[:msgLen], &msg)
	return
}
