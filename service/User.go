package service

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"imdemo/logic"
)

//用户数据
type User struct {
	Nickname   string
	Token      string
	OfflineMsg chan *Message
	conn       *websocket.Conn
}

//用户列表
var UserList = make(map[string]*User, 100)

//初始化用户
func CreateUser(nickname string, token string, conn *websocket.Conn) *User {
	//创建新用户
	if token == "" || UserList[token] == nil {
		token := logic.GetRandomString(8)
		user := &User{
			Nickname:   nickname,
			Token:      token,
			OfflineMsg: make(chan *Message, 1024),
			conn:       conn,
		}
		UserList[token] = user
		return user
	}
	//从map中获取用户数据
	user := UserList[token]
	user.conn = conn
	return user
}

//发送离线消息
func (u *User) SendOfflineMsg(mt int) bool {
	for offlineMsg := range u.OfflineMsg {
		data, _ := json.Marshal(offlineMsg)
		err := u.conn.WriteMessage(mt, data)
		if err != nil {
			return false
		}
	}
	return true
}

//发送消息
func (u *User) SendMsg(mt int, msg []byte) bool {
	err := u.conn.WriteMessage(mt, msg)
	if err != nil {
		offlineMsg := &Message{}
		_ = json.Unmarshal(msg, offlineMsg)
		u.OfflineMsg <- offlineMsg
	}
	return true
}

//群发消息
func SendAllMsg(mt int, msg []byte) bool {
	offlineMsg := &Message{}
	for _, user := range UserList {
		err := user.conn.WriteMessage(mt, msg)
		if err != nil {
			_ = json.Unmarshal(msg, offlineMsg)
			user.OfflineMsg <- offlineMsg
		}
	}
	return true
}
