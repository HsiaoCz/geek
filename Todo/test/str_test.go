package test

import (
	"fmt"
	"strings"
	"testing"
)

func TestContainsStr(t *testing.T) {
	str := "1233456@qq.com"
	ok := strings.Contains(str, "@")
	fmt.Println(ok)
}
