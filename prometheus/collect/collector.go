package collect

import "github.com/prometheus/client_golang/prometheus"

func NewDemoCollector() *DemoCollector {
	return &DemoCollector{
		queueLengthDesc: prometheus.NewDesc(
			"magedu_mcube_demo_queue_length",
			"The number of items in the queue",
			// 动态标签
			[]string{"instance_id", "instance_name"},
			// 静态追踪
			prometheus.Labels{"module": "http-server"},
		),
		// 动态的value列表，这里必须与声明的动态标签的key一一对应
		labelValues: []string{"mq_001", "kafka01"},
	}
}

type DemoCollector struct {
	queueLengthDesc *prometheus.Desc
	labelValues     []string
}

func (d *DemoCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- d.queueLengthDesc
}

func (c *DemoCollector) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(c.queueLengthDesc, prometheus.GaugeValue, 100, c.labelValues...)
}
