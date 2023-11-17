package codec

import (
	"bufio"
	"encoding/gob"
	"io"
	"log"
)

// 此目录用来对消息进行编解码
// 一个典型的rpc调用
// err=client.Call("Arith.Multiply",args,&reply)
// 客户端发送的请求包括服务名，方法名，参数
// 服务端的响应包含error，返回值reply

// 这里抽象请求和响应中的参数和返回值为一个结构体

type Header struct {
	ServiceMethod string // 格式化为"服务名.方法"
	Seq           uint64 // 请求的序号，或者id，用来区分不同的请求
	Error         string // 错误信息，如果服务端发生错误，将错误信息置于Error中
}

// 对消息体进行编解码的接口
type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(any) error
	Write(*Header, any) error
}

type NewCodecFunc func(io.ReadWriteCloser) Codec

type Type string

const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json" // not implemented
)

var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}

type GobCodec struct {
	conn io.ReadWriteCloser
	buf  *bufio.Writer
	dec  *gob.Decoder
	enc  *gob.Encoder
}

var _ Codec = (*GobCodec)(nil)

func NewGobCodec(conn io.ReadWriteCloser) Codec {
	buf := bufio.NewWriter(conn)
	return &GobCodec{
		conn: conn,
		buf:  buf,
		dec:  gob.NewDecoder(conn),
		enc:  gob.NewEncoder(buf),
	}
}

func (c *GobCodec) ReadHeader(h *Header) error {
	return c.dec.Decode(h)
}

func (c *GobCodec) ReadBody(body interface{}) error {
	return c.dec.Decode(body)
}

func (c *GobCodec) Write(h *Header, body interface{}) (err error) {
	defer func() {
		_ = c.buf.Flush()
		if err != nil {
			_ = c.Close()
		}
	}()
	if err := c.enc.Encode(h); err != nil {
		log.Println("rpc codec: gob error encoding header:", err)
		return err
	}
	if err := c.enc.Encode(body); err != nil {
		log.Println("rpc codec: gob error encoding body:", err)
		return err
	}
	return nil
}

func (c *GobCodec) Close() error {
	return c.conn.Close()
}
