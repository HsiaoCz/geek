package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Student struct {
	Username  string
	Identity  string
	Age       string
	Class     string
	Classmate string
}

func (s Student) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

func main() {
	// read config
	// 1 open the file
	fmt.Println("hello this is readetc")
	if err := OpenConfig("./stu.txt"); err != nil {
		log.Fatal(err)
	}
}

func OpenConfig(filename string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	read := bufio.NewReader(file)
	for {
		str, err := read.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(str)
		etcslice := strings.Split(string(str), " ")
		var stu Student
		if etcslice[0] != "Username" {
			stu = Student{
				Username:  etcslice[0],
				Identity:  etcslice[1],
				Age:       etcslice[2],
				Class:     etcslice[3],
				Classmate: etcslice[4],
			}
		}
		fmt.Println(stu)
	}
	return nil
}
