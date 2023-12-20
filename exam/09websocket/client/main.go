package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	c, _, err := websocket.Dial(ctx, "ws://localhost:3001/ws", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close(websocket.StatusInternalError, "内部错误")

	if err := wsjson.Write(ctx, c, "Hello WebSocket Server"); err != nil {
		log.Fatal(err)
	}
	var v interface{}
	if err := wsjson.Read(ctx, c, &v); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("接收到服务端响应:%v\n", v)
	c.Close(websocket.StatusNormalClosure, "")
}
