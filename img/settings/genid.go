package settings

import (
	"math/rand"
	"time"
)

// GenIdentity 生成用户id

func GenIdentity() int {
	randm := rand.New(rand.NewSource(time.Now().UnixNano()))
	return randm.Intn(100000000)
}
