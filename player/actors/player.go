package actors

import (
	"fmt"

	"github.com/anthdm/hollywood/actor"
)

type Player struct{}

func NewPlayer() actor.Receiver {
	return &Player{}
}

func (p *Player) Receive(c *actor.Context) {
	switch c.Message().(type) {
	case actor.Started:
		fmt.Println("player started")
	}
}
