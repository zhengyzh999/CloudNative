package main

import (
	"CloudNative/pprof/data"
	"CloudNative/pprof/data/block"
	"CloudNative/pprof/data/cpu"
	"CloudNative/pprof/data/goroutine"
	"CloudNative/pprof/data/mem"
	"CloudNative/pprof/data/mutex"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
	"time"
)

var cmds = []data.Cmd{
	&cpu.Cpu{},
	&mem.Mem{},
	&block.Block{},
	&goroutine.Goroutine{},
	&mutex.Mutex{},
}

func main() {
	log.SetFlags(log.Llongfile)
	log.SetOutput(os.Stdout)
	// 开启对锁调用的跟踪
	runtime.SetMutexProfileFraction(1)
	// 开启对阻塞作用的跟踪
	runtime.SetBlockProfileRate(1)

	cpuFile, err := os.OpenFile("out/cpu.out", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// 开始采集
	err = pprof.StartCPUProfile(cpuFile)
	// 停止采集
	defer pprof.StopCPUProfile()
	defer cpuFile.Close()

	memFile, err := os.OpenFile("out/mem.out", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// 开始采集
	err = pprof.WriteHeapProfile(memFile)
	defer memFile.Close()
	// 业务代码

	traceFile, err := os.OpenFile("out/trace.out", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// 开始采集
	err = trace.Start(traceFile)
	// 关闭采集
	defer trace.Stop()
	defer traceFile.Close()
	// 业务代码

	go func() {
		// http监听是个阻塞操作，放到goroutine中，阻塞单个协程，不阻塞主协程
		http.ListenAndServe("localhost:6060", nil)
	}()
	for i := 0; i < 10; i++ {
		for _, cmd := range cmds {
			cmd.Run()
		}
		time.Sleep(time.Second)
	}
}
