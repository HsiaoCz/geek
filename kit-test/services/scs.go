package services

import (
	"context"
	"errors"
	"math"
)

type AddServiceI interface {
	Sum(ctx context.Context, a int, b int) (int, error)
	Concat(ctx context.Context, a string, b string) (string, error)
}

type AddS struct{}

const (
	maxLen = 10
)

var (
	// ErrTwoZeroes  Sum方法的业务规则不能对两个0求和
	ErrTwoZeroes = errors.New("can't sum two zeroes")

	// ErrIntOverflow Sum参数越界
	ErrIntOverflow = errors.New("integer overflow")

	// ErrTwoEmptyStrings Concat方法业务规则规定参数不能是两个空字符串.
	ErrTwoEmptyStrings = errors.New("can't concat two empty strings")

	// ErrMaxSizeExceeded Concat方法的参数超出范围
	ErrMaxSizeExceeded = errors.New("result exceeds maximum size")
)

// Sum 对两个数字求和，实现AddService。
func (s *AddS) Sum(_ context.Context, a, b int) (int, error) {
	if a == 0 && b == 0 {
		return 0, ErrTwoZeroes
	}
	if (b > 0 && a > (math.MaxInt-b)) || (b < 0 && a < (math.MinInt-b)) {
		return 0, ErrIntOverflow
	}
	return a + b, nil
}

// Concat 连接两个字符串，实现AddService。
func (s *AddS) Concat(_ context.Context, a, b string) (string, error) {
	if a == "" && b == "" {
		return "", ErrTwoEmptyStrings
	}
	if len(a)+len(b) > maxLen {
		return "", ErrMaxSizeExceeded
	}
	return a + b, nil
}
