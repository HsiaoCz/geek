package settings

import (
	"math/rand"
	"time"
)

func GenIdentity() int {
	mrand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return mrand.Intn(100000000)
}
