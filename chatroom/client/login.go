package main

// //用于登录校验
// func login(userId int, userPwd string) (err error) {
// 	//定义协议
// 	// fmt.Printf("userid = %d ,userpwd = %s \n", userId, userPwd)
// 	// return nil

// 	//暂时写死，应该是读配置
// 	conn, err := net.Dial("tcp", "localhost:8889")
// 	if err != nil {
// 		fmt.Println("net.Dial error:", err)
// 		return
// 	}
// 	defer conn.Close() //延时关闭
// 	//准备消息体：登录信息包装成loginMessage的消息体
// 	var mes common.Message
// 	//①准备login的type数据
// 	mes.Type = common.LoginMessageType
// 	//②准备login的data数据
// 	var loginMes = common.LoginMessage{
// 		UserId:  userId,
// 		UserPwd: userPwd,
// 	}
// 	//②-1 需要序列化成string
// 	data, err := json.Marshal(loginMes)
// 	if err != nil {
// 		fmt.Println("json.Marshal(loginMessage.Data) error:", err)
// 		return
// 	}
// 	mes.Data = string(data)

// 	// 发送消息给服务端
// 	//1. 将message消息体进行序列化
// 	msg, err := json.Marshal(mes)
// 	if err != nil {
// 		fmt.Println("json.Marshal(loginMessage) error:", err)
// 		return
// 	}
// 	//2.先发送消息体的data长度
// 	var pkgLen uint32 = uint32(len(msg)) //data的长度还是msg的长度?
// 	var buffer [4]byte
// 	binary.BigEndian.PutUint32(buffer[:4], pkgLen)

// 	//3.正式开始发送
// 	n, err := conn.Write(buffer[:4])
// 	if n != 4 || err != nil {
// 		fmt.Printf("conn.Write失败,长度%d,err:%v\n", n, err)
// 		return
// 	}

// 	fmt.Printf("客户端，发送的长度:%d,内容=%s\n", pkgLen, string(msg))

// 	return nil
// }
