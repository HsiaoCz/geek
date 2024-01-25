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

3 直方图: 主要用于统计指标值的一个分布情况

使用 Bucket:设置横轴区间，只设置上限，不设置下限
0~100
0~90
0~80
0~70
0~60

直方图，统计请求耗时分布
0~100ms 请求个数
0~500ms 请求个数

```go
requestDurations:=prometheus.NewHistogram(prometheus.HistogramOpts{
    Name:"http_request_duration_secondes",
    Help:"A histogram of HTTP request duration in seconds. "
    // bucket  配置，第一个bucket包括所有在0.5s内完成的请求，最后一个包含所有在10s内完成的请求
    Buckets: []float64{0.05,0.1,0.25,1,2.5,5,10}
})
```

示例

```go
func TestHistogram(t *testing.T){
 requestDurations:=prometheus.NewHistogram(prometheus.HistogramOpts{
    Name:"http_request_duration_secondes",
    Help:"A histogram of HTTP request duration in seconds. "
    // bucket  配置，第一个bucket包括所有在0.5s内完成的请求，最后一个包含所有在10s内完成的请求
    Buckets: []float64{0.05,0.1,0.25,1,2.5,5,10}
})

// 添加值
for _,v:=range []float64{0.01,0.02,0.3,0.4,0.6,0.7,5.5,11}{
    requestDurations.Observe(v)
}
}
```

分位数 summaries

这种类型的指标，就是用于计算分位数的，因此他需要配置一个核心参数，你需要统计那个百分位用 NewSummary 来构建该类指标

```go
requestDurations:=prometheus.NewSummary(prometheus.SummaryOpts{
    Name:"http_request_duration_secondes",
    Help:"A summary of the HTTP request durations in seconds",
    Objectives:map[float64]float64{
        0.5:0.05, // 第五十个百分位，最大绝对误差为0.05
        0.9:0.01, // 第90个百分位，最大绝对误差为0.01
        0.99:0.001, // 第99个百分位，最大绝对误差为0.01
    }
})
```

指标标签:

标签分为 2 类：

- 静态标签:constLabels，在指标创建时，就提前声明好，采集过程中永不变动的
- 动态标签:variableLabels，用于在指标的收集过程中动态的补充标签，比如 kafka 集群的 exporter 需要动态补充 instance_id

要让你的指标支持动态标签，有专门的构造函数,对应关系；

- NewGauge() 变成 NewGaugeVec()
  以此类推

指标注册

```go
prometheus 定义了一个注册表的接口

// 指标注册接口
type Registerer interface{
    // 注册采集器，有异常会报错
    Register(Collector)error
    // 注册采集器，有异常会panic
    MustRegister(...Collector)
    // 注销该采集器
    Unregister(collector)bool
}

// prometheus实现了一个默认的register对象
var (
    defaultRegistry    = NewRegistry()
    DefaultRegisterer Registerer= defaultRegistry
    DefaultGatherer  Gatherer = defaultRegistry
)

我们通过prometheus提供的MustRegister可以将我们自定义指标注册进去

// 在默认的注册表中注册该指标
prometheus.MustRegister(temp)
prometheus.Register()
peometheus.Unregister()
```

自定义注册表

- 使用 NewRegistry()创建一个全新的注册表
- 通过注册表对象的 MustRegister 把指标注册到自定义的注册表中

暴露指标的时候必须通过调用 promehttp.HandleFor()函数来创建一个专门针对我们自定义注册表的 HTTP 处理器，我们还需要 promehttp.HandleOpts 配置对象的 Registry 字段中传递我们的注册表对象

```go
// 暴露指标
http.Handle("/metrics",promhttp.HandlerFor(registry,promhttp.HandlerOpts{Registry:registry}))
http.ListenAndServe(":8050",nil)
```

其实prometheus在客户端中默认有如下Collector供我们选择
只需要把我们需要的添加到我们自定义的注册表中即可

```go
// 添加process 和go运行时指标到我们自定义的注册表中
registry.MustRegister(peometheus.NewProcessCollector(prometheus.ProcessCollectionOpts{}))
registry.MustRegister(prometheus.NewGoCollector())
```

Collector接口解读

```go
type Collector interface{
    // 指标的一些描述信息，就是#标识的部分
    // 注意这里使用的时指针，因为描述信息，全局存储一份即可
    Describe(chan<- *Desc)
    // 指标的数据，比如promehttp_metric_handler_errors_total{cause="gathering"} 0
    // 这里没有使用指针，因为每次采集的值都是独立的
    Collect(chan<- Metric)
}
```