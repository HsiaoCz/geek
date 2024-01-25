package controller

import "github.com/prometheus/client_golang/prometheus"

type RocketMQCollector struct{}

func (c *RocketMQCollector) Describe(chan<- *prometheus.Desc) {}
func (c *RocketMQCollector) Collect(chan<- prometheus.Metric) {}
