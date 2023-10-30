package utils

import (
	"math/rand"
	"time"
)

func GenIdentity() int64 {
	mrand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return mrand.Int63n(1000000000)
}
