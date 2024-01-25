package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(str) != 0 {
				fmt.Println(str)
			}
			break
		}
		if err != nil {
			return
		}
		fmt.Println(str)
	}
}
