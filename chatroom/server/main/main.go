package main

import (
	"fmt"
	"net"
	"test/chatroom/server/model"
	"test/chatroom/server/processor"
	"time"

	"github.com/garyburd/redigo/redis"
)

var RedisPool *redis.Pool

func init() {
	RedisPool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
	model.MyUserDao = model.NewUserDao(RedisPool)
}

func main() {

	fmt.Println("服务器在8889端口开始监听。。。")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.listen error:", err)
		return
	}

	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("lisen accept error:", err)
		}
		//连接成功够，启动协程与客户端进行通讯
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()

	//创建一个总控
	processor := &processor.Processor{
		Conn: conn,
	}
	processor.Process()
}
