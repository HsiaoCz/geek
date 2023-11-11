package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

var (
	// 新用户到来，通过该channel进行登记
	enteringChannel = make(chan *User)
	// 用户离开，通过该channel 进行登记
	leavingChannel = make(chan *User)
	// 广播专用的用户普通消息channel，缓冲是尽可能避免出现异常情况堵塞
	messageChannel = make(chan string, 8)
)

type User struct {
	ID   int
	Addr string
	// 进入的时间
	EnterAt time.Time
	// 当前用户发送消息的通道
	MessageChannel chan string
}

func (u *User) String() string {
	return <-u.MessageChannel
}

// 基于tcp的聊天室
func main() {
	listen, err := net.Listen("tcp", ":2020")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

// 记录聊天室用户，并进行消息广播
// 1.新用户进来
// 2.用户普通消息
// 3.用户离开
func broadcaster() {
	users := make(map[*User]struct{})
	for {
		select {
		case user := <-enteringChannel:
			// 新用户进入
			users[user] = struct{}{}
		case user := <-leavingChannel:
			// 用户离开
			delete(users, user)
			// 避免groutine泄露
			close(user.MessageChannel)
		case msg := <-messageChannel:
			// 给所有在线用户发送消息
			for user := range users {
				user.MessageChannel <- msg
			}
		}
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	// 1.新用户进来
	user := &User{
		ID:             GenUserID(),
		Addr:           conn.RemoteAddr().String(),
		EnterAt:        time.Now(),
		MessageChannel: make(chan string, 8),
	}
	// 2.当前一个新的goroutine中，用来进行读操作，因此需要开一个
	// groutine用来写操作
	// 读写goroutine之间可以通过channel进行通信
	go sendMessage(conn, user.MessageChannel)

	// 3.给当前用户发送欢迎消息：给所有用户告知新用户到来
	user.MessageChannel <- "Welcome," + user.String()
	messageChannel <- "user:`" + strconv.Itoa(user.ID) + "`has enter"

	// 4.将该记录到全局的用户列表中，避免使用锁
	enteringChannel <- user

	// 5. 循环读取用户的输入
	input := bufio.NewScanner(conn)
	for input.Scan() {
		messageChannel <- strconv.Itoa(user.ID) + ":" + input.Text()
	}
	if err := input.Err(); err != nil {
		log.Println("读取错误:", err)
	}

	// 6.用户离开
	leavingChannel <- user
	messageChannel <- "user:`" + strconv.Itoa(user.ID) + "`has left"
}

func GenUserID() int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100000)
}

func sendMessage(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
