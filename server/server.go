package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"imdemo/global"
	"log"
	"net/http"
	"strconv"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var uid = 0;

//用户连接信息
var usersMap = make(map[int] *global.User)

var msg global.Msg

func main() {
	fmt.Printf("webSocket Server listen at 9990\n")
	http.HandleFunc("/ws", WebSocketHandleFunc)
	log.Fatal(http.ListenAndServe("localhost:9990", nil))
}

func WebSocketHandleFunc(w http.ResponseWriter, r *http.Request)  {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	uid += 1
	usersMap[uid] = &global.User{
		uid,
		c,
	}
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		_ = json.Unmarshal(message, &msg)
		switch msg.MsgType {
		case 1:
			//私发
			user := usersMap[msg.Uid]
			_ = user.Conn.WriteMessage(mt, []byte("收到私发红包" + strconv.FormatFloat(float64(msg.Money), 'f', 2, 64)))
		case 2:
			//群发
			for _, user := range usersMap{
				_ = user.Conn.WriteMessage(mt, []byte("收到群发红包" + strconv.FormatFloat(float64(msg.Money), 'f', 2, 64)))
			}
		case 3:
			//拆红包
		}
		log.Printf("recv: %s\n", message)
		err = c.WriteMessage(mt, []byte("发送成功"))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}