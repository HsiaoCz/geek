package main

import (
	"log"

	"github.com/HsiaoCz/geek/player/actors"
	"github.com/anthdm/hollywood/actor"
	"github.com/anthdm/hollywood/cluster"
	"github.com/anthdm/hollywood/remote"
)

func main() {
	bootstripAddr := cluster.MemberAddr{
		ListenAddr: "127.0.0.1:3001",
		ID:         "A",
	}
	r := remote.New(":3002", nil)
	e, err := actor.NewEngine(&actor.EngineConfig{Remote: r})
	if err != nil {
		log.Fatal(err)
	}
	c, err := cluster.New(cluster.Config{
		ID:              "B",
		Engine:          e,
		Region:          "us-west",
		ClusterProvider: cluster.NewSelfManagedProvider(bootstripAddr),
	})
	if err != nil {
		log.Fatal(err)
	}
	c.RegisterKind("player", actors.NewPlayer, nil)
	c.Start()
	select {}
}
