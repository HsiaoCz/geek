prometheus exporter

数据格式：

- 通讯协议

  - HTTP 协议
  - 服务端实现了 gzip

- 数据格式
  - text/plain：文本协议

数据格式

prometheus 是拉取数据的监控模型，它对客户端暴露的数据格式要求如下:
(指标描述)

# HELP go_goroutines Number of goroutines that currency exist.

# TYPE go_goroutines gauge

(指标名称)go_goroutine 19 指标的值

徒手开发一个 exporter,只要满足 prometheus 要求的数据规范就可以被 prometheus

exporter 指标类型:

1. Gauges 仪表盘
2. Counters 计数器
3. Histograms 直方图
4. Summaries 摘要

1 Gauges 仪表盘
最常见的指标，也就是实时指标，值是什么就返回什么，并不会进行加工处理
SDK 提供了该指标的构造函数:NewGauge

```go
queueLength:=prometheus.NewGague(prometheus.GaugeOpts{
    // Namespace,Subsystem,Name 会拼接成指标的名称:magedu_mcube_dmo_queue_length
    // 其中Name 是必填参数
    Namespace:"magedu",
    Subsystem:"mcube_demo",
    // 指标的描述信息
    Help: "The number of items in the queue .",
    // 指标标签
    ConstLabels:map[string]string{
        "module":"http-server",
    },
})
```

看一个具体的例子:

```go
func TestGauge(t *testing.T){
    queueLength:=prometheus.NewGague(prometheus.GaugeOpts{
    // Namespace,Subsystem,Name 会拼接成指标的名称:magedu_mcube_dmo_queue_length
    // 其中Name 是必填参数
    Namespace:"magedu",
    Subsystem:"mcube_demo",
    Name:"queue_length",
    // 指标的描述信息
    Help: "The number of items in the queue .",
    // 指标标签
    ConstLabels:map[string]string{
        "module":"http-server",
    },
})
}

queueLength.Set(100)
```

最后得到的就是这样的一个东西:

# HELP magedu_mcube_demo_queue_length The number of items in the queue.

# TYPE magedu_mcube_demo_queue_length gauge

magedu_mcube_demo_queue_length{module="http-server"} 100

2 Counter 是计算指标，用于统计次数使用，通过 prometheus.NewCounter()函数来初始化指标对象

```go
totalRequests := prometheus.NewCounter(prometheus.CounterOpts{
    Name:"http_requests_total",
    Help:"The total number of handled HTTP requests",
})
```

```go
func TestCounter(t *testing.T){
 totalRequests := prometheus.NewCounter(prometheus.CounterOpts{
    Name:"http_requests_total",
    Help:"The total number of handled HTTP requests",
})
  for i:=0;i<10;i++{
    totalRequests.Inc()
  }
}
```

3  直方图: 主要用于统计指标值的一个分布情况

使用Bucket:设置横轴区间，只设置上限，不设置下限
0~100
0~90
0~80
0~70
0~60