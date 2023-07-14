## context

context 包的核心 API 有四个：

- context.WithValue:设置键值对，并且返回一个新的 context
- context.WithCancel
- context.WithDeadline
- context.WithTimeout:三者都返回一个可取消的 context 实例，和取消函数

context 为了保证线程安全
context 实例是不可变的，每次创建都是新建

Context 接口的核心 API 有四个

- Deadline：返回过期时间，如果 ok 为 false，说明没有设置过期时间
- Done:返回一个 channel,一般用于监听 context 实例的信号，比如说过期，或者正常关闭
- Err:返回一个错误用于表达 context 发送了什么 canceled 正常关闭，deadlineExceeded 过期超时
- value:取值，非常常用

context包我们主要用来做两件事情：
- 安全传递数据
- 控制链路

安全传递数据，是指在请求执行上下文中线程安全地传递数据，依赖于WithWValue方法

context的实例之间存在父子关系
- 当父亲取消或者超时时，所有派生的子context都被取消或者超时
- 当找key的时候，子context先看自己有没有，没有则去祖先里面找

控制是从上至下的，查找则是从下至上的