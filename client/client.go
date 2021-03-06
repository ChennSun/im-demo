package main

import (
	"fmt"
)

//消息体

type Msg struct {
	msgType int
	money   float32
	uid     int
}

var c = make(chan *Msg, 1024)

func main() {
	go getIn()
	go scanIn()
	for  {
		
	}
}

func getIn()  {
	for  {
		a := <-c
		fmt.Print(a.money)
	}
}

func scanIn() {
	for {
		var (
			money          float32
			red_money_name string
			red_money_type int
		)
		fmt.Print("请输入发送红包金额：")
		fmt.Scanln(&money)
		fmt.Printf("红包金额为: %f\n", money)
		fmt.Print("请选择红包类型(1.私发 2.群发)")
		fmt.Scanln(&red_money_type)
		switch red_money_type {
		case 1:
			red_money_name = "私发"
		case 2:
			red_money_name = "群发"
		}
		msg := &Msg{
			msgType: red_money_type,
			money: money,
			uid: 100,
		}
		c <- msg
		fmt.Printf("红包类型为: %s\n", red_money_name)
	}
}
