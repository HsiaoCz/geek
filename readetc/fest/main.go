package main

import (
	"bufio"
	"encoding/json"
	"fmt"
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
	some, err := GetData("./stu.txt")
	if err != nil {
		log.Fatal(err)
	}
	var stu []Student
	for i := 0; i < len(some)/5; i++ {
		stu = append(stu, Student{
			Username:  some[0+5*i],
			Identity:  some[1+5*i],
			Age:       some[2+5*i],
			Class:     some[3+5*i],
			Classmate: some[4+5*i],
		})
	}
	fmt.Println(stu)
}

func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// 存储读取到的文件的缓冲区
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func GetData(filename string) ([]string, error) {
	// 读取文件
	etcstu, err := ReadLines(filename)
	if err != nil {
		return nil, err
	}
	// 将文件内容转换为结构体的值
	// 这里把第一行去掉
	var etcs []string
	for i := 1; i < len(etcstu); i++ {
		etcstr := etcstu[i]
		str := strings.ReplaceAll(etcstr, " ", "/")
		ss := strings.Split(str, "/")
		for j := 0; j < len(ss); j++ {
			if ss[j] == "" {
				continue
			}
			etcs = append(etcs, ss[j])
		}
	}
	return etcs, nil
}
