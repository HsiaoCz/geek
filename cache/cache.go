package cache

// ByteView用来表示缓存值
type ByteView struct {
	b []byte
}

// Len 返回ByteView的长度
func (v ByteView) Len() int {
	return len(v.b)
}

// ByteSlice 返回复制的ByteView内部的切片
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

// String 对切片进行类型转换
func (v ByteView) String() string {
	return string(v.b)
}

// cloneBytes 对ByteView内的b的值进行复制
// 这里的作用在于b是只读的，返回一个拷贝，避免被外部的程序修改
func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}