package data_test

import (
	"CloudNative/pprof/data/block"
	"CloudNative/pprof/data/cpu"
	"CloudNative/pprof/data/goroutine"
	"CloudNative/pprof/data/mem"
	"CloudNative/pprof/data/mutex"
	"testing"
)

// 基准测试
func BenchmarkData(b *testing.B) {
	b.Run("block", func(b *testing.B) {
		o := block.Block{}
		for i := 0; i < b.N; i++ {
			o.Run()
		}
	})
	b.Run("cpu", func(b *testing.B) {
		o := cpu.Cpu{}
		for i := 0; i < b.N; i++ {
			o.Run()
		}
	})
	b.Run("mem", func(b *testing.B) {
		o := mem.Mem{}
		for i := 0; i < b.N; i++ {
			o.Run()
		}
	})
	b.Run("goroutine", func(b *testing.B) {
		o := goroutine.Goroutine{}
		for i := 0; i < b.N; i++ {
			o.Run()
		}
	})
	b.Run("mutex", func(b *testing.B) {
		o := mutex.Mutex{}
		for i := 0; i < b.N; i++ {
			o.Run()
		}
	})
}
