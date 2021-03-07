package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("webSocket Server listen at 9990\n")
	http.HandleFunc("/ws", WebSocketHandleFunc)
	log.Fatal(http.ListenAndServe("localhost:9990", nil))
}
