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

	go func() {
		// http监听是个阻塞操作，放到goroutine中，阻塞单个协程，不阻塞主协程
		http.ListenAndServe("localhost:6060", nil)
	}()
	for {
		for _, cmd := range cmds {
			cmd.Run()
		}
		time.Sleep(time.Second)
	}
}
