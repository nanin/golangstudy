package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("服务器在8889端口开始监听。。。")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.listen error:", err)
		return
	}

	defer listen.Close()

	for {
		fmt.Println("等待客户端来连接。。。")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("lisen accept error:", err)
		}
		//连接成功够，启动协程与客户端进行通讯
		go process(conn)
	}
}

func process(conn net.Conn) {
	//延时关闭
	defer conn.Close()

	//todo:客户端发过来的消息处理
	for {
		buf := make([]byte, 8096)
		n, err := conn.Read(buf[0:4])
		if n != 4 && err != nil {
			fmt.Printf("conn.Read,长度%d,err:%v\n", n, err)
			return
		}
		fmt.Println("读到buffer=", buf[:4])
	}
}
