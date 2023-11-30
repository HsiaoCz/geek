package main

import (
	"log"

	"github.com/go-playground/validator/v10"
)

// go 使用validator进行参数校验
type Test struct {
	Field1 string `validate:"required"`
	Email  string `validate:"email"`
}

// 实例化验证对象
var validate = validator.New()

func main() {
	test := &Test{
		Field1: "hello",
		Email:  "1222",
	}
	if err := validate.Struct(test); err != nil {
		log.Fatal(err)
	}
}
