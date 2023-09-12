package queue

import (
	"errors"
	"sync"

	"github.com/HsiaoCz/geek/middle/model"
)

// 这里实现一个队列
// 用户发帖后，会放入队列里面
// 管理员可以从队列里面取文章进行人工审核
// 人工审核完成后，可以提交发布，也就是可以进入数据库
// 否则会退回给用户，只有用户删除之后，才会真正的删除文章

// ArtQueue  这个队列只能存article这种数据
// 添加一个lock字段，避免并发执行的竞态问题

// 1、当队列里面没有数据的时候，需要返回错误
// 2、需要给队列一个最大长度限制，超过这个限制就不能入队了，这里也要思考一下，如果数据不能入队了怎么办？
// 直接返回吗？
type ArtQueue struct {
	data   []model.Article
	lock   sync.Mutex
	length int
}

// InQueue 入队
func (a *ArtQueue) InQueue(article model.Article) error {
	a.lock.Lock()
	defer a.lock.Unlock()
	if len(a.data) >= a.length {
		return errors.New("数据已满")
	}
	a.data = append(a.data, article)
	return nil
}

// OutQueue 出队
func (a *ArtQueue) OutQueue() (v model.Article, err error) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if len(a.data) == 0 {
		return model.Article{}, errors.New("队列为空")
	}
	article := a.data[0]
	a.data = a.data[1:]
	return article, nil
}

// 这里的长度可以通过配置文件来设置
func NewArtQueue(length int) *ArtQueue {
	return &ArtQueue{
		data:   make([]model.Article, 0),
		length: length,
	}
}
