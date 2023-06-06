package common

const (
	LoginMessageType       = "LoginMessage"
	LoginResultMessageType = "LoginResultMessage"
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息内容
}

type LoginMessage struct {
	UserId  int    `json:"userid"`
	UserPwd string `json:"userpwd"`
}

type LoginResultMessage struct {
	Code  int    `json:"code"`  //返回状态码：500该用户未注册、200登录成功
	Error string `json:"error"` //返回错误信息
}
