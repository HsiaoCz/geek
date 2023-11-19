package main

import (
	"log"

	"github.com/HsiaoCz/geek/exam/04word/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
