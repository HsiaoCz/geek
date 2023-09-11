package queue

import (
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
type ArtQueue struct {
	data []model.Article
	lock sync.Mutex
}

// InQueue 入队
func (a *ArtQueue) InQueue(article model.Article) {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.data = append(a.data, article)
}

// OutQueue 出队
func (a *ArtQueue) OutQueue() model.Article {
	a.lock.Lock()
	defer a.lock.Unlock()
	article := a.data[0]
	a.data = a.data[1:]
	return article
}
