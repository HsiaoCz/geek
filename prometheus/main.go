package main

import (
	"net/http"

	collectioner "github.com/HsiaoCz/geek/prometheus/collect"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// 创建一个自定义的注册表
	registry := prometheus.NewRegistry()

	// 可选：添加process 和Go的运行时指标到我们自定义的注册表中
	registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	registry.MustRegister(collectors.NewGoCollector())

	// 注册自定义采集器
	registry.MustRegister(collectioner.NewDemoCollector())
	// 暴露指标
	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{Registry: registry}))
	http.ListenAndServe(":8050", nil)
}
