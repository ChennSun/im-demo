package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"imdemo/service"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	//不进行同源检测
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocketHandleFunc(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	token := r.FormValue("token")
	nickname := r.FormValue("nickname")
	user := service.CreateUser(nickname, token, c)
	_ = c.WriteMessage(1, []byte("你的token:"+user.Token))
	go user.SendOfflineMsg(1)
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		msg := &service.Message{}
		_ = json.Unmarshal(message, &msg)
		switch msg.Type {
		case 1:
			//私发
			user := service.UserList[msg.ToToken]
			_ = user.SendMsg(mt, message)
		case 2:
			//群发
			service.SendAllMsg(mt, message)
		}
		log.Printf("recv: %s\n", message)
		err = c.WriteMessage(mt, []byte("发送成功"))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
