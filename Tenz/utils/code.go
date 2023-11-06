package utils

import (
	"math/rand"
	"time"
)

// 发送验证码
func SendCode(phone_num int) error {
	// 这里先生成一个六位数的验证码
	randt := rand.New(rand.NewSource(time.Now().UnixNano()))
	_ = randt.Intn(1000000)
	// 将六位数的发送到用户的手机上
	// 将生成的验证码保存到redis上并设置60s过期
	// redis的key为phone_num:code
	// 然后将验证码返回
	// 不过这里还应该再返回一个错误
	// 因为手机号可能无法发送验证码
	return nil
}

// 验证验证码
func ParseCode(code int, phone_num int) bool {
	// 对拿到的验证码进行验证
	// 对比一致就说明验证通过
	return true
}
