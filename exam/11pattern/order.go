package pattern

import "fmt"

// 命令模式
// 将一个请求封装成一个对象，从而让我们可以用不同的请求对客户端进行参数化
// 命令模式可以将请求发送者和接收者完全解耦，发送者和接收者之间没有直接的引用关系
// 发送请求的对象只需知道如何发送请求，而不必知道如何完成请求

type Cooker struct{}

func (c *Cooker) MakeChuaner() {
	fmt.Println("师傅烤了羊肉串...")
}

func (c *Cooker) MakeJiChi() {
	fmt.Println("师傅烤了鸡翅...")
}

// 抽象的命令
type Command interface {
	Make()
}

type CommandCookChicken struct {
	cooker *Cooker
}

func (cmd *CommandCookChicken) Make() {
	cmd.cooker.MakeJiChi()
}

type CommandCookChuaner struct {
	cooker *Cooker
}

func (cmd *CommandCookChuaner) Make() {
	cmd.cooker.MakeChuaner()
}

// 服务员，命令的收集者
type Waiter struct {
	CmdList []Command
}

func (w Waiter) Notify() {
	if w.CmdList == nil {
		return
	}
	for _, cmd := range w.CmdList {
		cmd.Make()
	}
}
