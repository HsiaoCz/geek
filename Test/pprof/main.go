package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

// pprof 性能调优
// 性能调优主要考虑这几个
// cpu,memory,groutine的情况
// go语言内置的获取程序运行信息的标准库有两个
// runtime/pprof：采集工具型应用运行数据进行分析
// net/http/pprof：采集服务型应用运行时数据进行分析

// 开启pprof
// 每隔一段时间（10ms）就会收集下当前的堆栈信息
// 获取各个函数占用的CPU以及内存资源；
// 最后通过对这些采样数据进行分析，形成一个性能分析报告。

// 如果你的应用程序是运行一段时间就结束退出类型
// 这种情况下可以使用runtine/pprof

// cpu性能分析
// 开启cpu性能分析
// pprof.StartCPUProfile(w io.Writer)
// 关闭cpu性能分析
// pprof.StopCPUProfile()

// memory的性能分析
// pprof.WriteHeapProfile(w io.Writer)
// go tool pprof默认是使用-inuse_space进行统计
// 还可以使用-inuse-objects查看分配对象的数量

// 服务型应用的性能分析
// 使用默认的http.DefaultServeMux()
// 只需要按照这种方式 import _ "net/http/pprof" 导入即可

// 如果使用自定义的Mux
// r.HandleFunc("/debug/pprof/", pprof.Index)
// r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
// r.HandleFunc("/debug/pprof/profile", pprof.Profile)
// r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
// r.HandleFunc("/debug/pprof/trace", pprof.Trace)

// 使用gin框架，可以使用github.com/gin-contrib/pprof
// 可以访问/debug/pprof/
// /debug/pprof/profile：访问这个链接会自动进行 CPU profiling，持续 30s，并生成一个文件供下载
// /debug/pprof/heap： Memory Profiling 的路径，访问这个链接会得到一个内存 Profiling 结果的文件
// /debug/pprof/block：block Profiling 的路径
// /debug/pprof/goroutines：运行的 goroutines 列表，以及调用关系

// go tool pprof命令
// go tool pprof [binary] [source]
// binary 是应用的二进制文件，用来解析各种符号；
// source 表示 profile 数据的来源，可以是本地的文件，也可以是 http 地址。

func logicCode() {
	var c chan int
	for {
		select {
		case v := <-c:
			fmt.Printf("recv from chan, value:%v\n", v)
		default:
			fmt.Println("hello")
		}
	}
}

func main() {
	var isCPUPprof bool
	var isMemPprof bool

	flag.BoolVar(&isCPUPprof, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMemPprof, "mem", false, "turn mem pprof on")
	flag.Parse()

	if isCPUPprof {
		file, err := os.Create("./cpu.pprof")
		if err != nil {
			fmt.Printf("create cpu pprof failed, err:%v\n", err)
			return
		}
		pprof.StartCPUProfile(file)
		defer pprof.StopCPUProfile()
	}
	for i := 0; i < 8; i++ {
		go logicCode()
	}
	time.Sleep(20 * time.Second)
	if isMemPprof {
		file, err := os.Create("./mem.pprof")
		if err != nil {
			fmt.Printf("create mem pprof failed, err:%v\n", err)
			return
		}
		pprof.WriteHeapProfile(file)
		file.Close()
	}
}

// 可以将代码编译之后，使用./main -cpu -mem 开启
// flat：当前函数占用CPU的耗时
// flat：:当前函数占用CPU的耗时百分比
// sun%：函数占用CPU的耗时累计百分比
// cum：当前函数加上调用当前函数的函数占用CPU的总耗时
// cum%：当前函数加上调用当前函数的函数占用CPU的总耗时百分比
// 最后一列：函数名称

// pprof 还可以和go test 的benchmark结合
// -cpuprofile：cpu profiling 数据要保存的文件地址
// -memprofile：memory profiling 数据要报文的文件地址

// 执行测试的同时，也会执行 CPU profiling，并把结果保存在 cpu.prof 文件中
// go test -bench . -cpuprofile=cpu.prof 
// 执行测试的同时，也会执行 Mem profiling，并把结果保存在 cpu.prof 文件中：
// go test -bench . -memprofile=./mem.prof