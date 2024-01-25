package main

import (
	"fmt"
	"log"

	"github.com/anthdm/hollywood/actor"
	"github.com/anthdm/hollywood/cluster"
	"github.com/anthdm/hollywood/remote"
)

func main() {
	r := remote.New(":3001", nil)
	e, err := actor.NewEngine(&actor.EngineConfig{Remote: r})
	if err != nil {
		log.Fatal(err)
	}
	c, err := cluster.New(cluster.Config{
		ID:              "A",
		Engine:          e,
		Region:          "eu-west",
		ClusterProvider: cluster.NewSelfManagedProvider(),
	})
	if err != nil {
		log.Fatal(err)
	}
	c.Engine().SpawnFunc(func(ctx *actor.Context) {
		switch msg := ctx.Message().(type) {
		case cluster.MemberJoinEvent:
			if msg.Member.ID == "B" {
				pid := c.Activate("player", &cluster.ActivationConfig{ID: "bob"})
				fmt.Println(pid)
			}
		}
	}, "event")
	c.Start()
	c.Activate("player", &cluster.ActivationConfig{ID: "bob"})
	select {}
}
