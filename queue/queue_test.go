package queue

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestQueueTimeOut(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	var q Queue[any]
	q.QPush(ctx, "1233")
	select {
	case <-ctx.Done():
		fmt.Println("取消")
	default:
		fmt.Println("hello")
	}
	cancel()
}
