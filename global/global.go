package global

import "github.com/gorilla/websocket"

//消息体
type Msg struct {
	MsgType int
	Money   float32
	Uid     int
}

//用户数据
type User struct {
	Uid int
	Conn *websocket.Conn
}

