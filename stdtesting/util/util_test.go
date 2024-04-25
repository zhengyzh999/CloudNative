package util

import (
	"fmt"
	"testing"
	"time"
)

var commonTestData []commonStruct

type commonStruct struct {
	Group         string
	SizeStr       string
	ExpectSize    int64
	ExpectSizeStr string
}

// 功能测试
func TestParseSize(t *testing.T) {
	testData := commonTestData
	for _, data := range testData {
		fmt.Println(data)
		size, sizeStr := ParseSize(data.SizeStr)
		if size != data.ExpectSize || sizeStr != data.ExpectSizeStr {
			t.Errorf("测试结果不符合预期: %+v", data)
		}
	}
}

// 功能测试子测试、并发测试
func TestParseSizeSub(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过测试用例 TestParseSizeSub")
	}
	testData := make(map[string][]commonStruct)
	for _, item := range commonTestData {
		group := item.Group
		_, ok := testData[group]
		if !ok {
			testData[group] = make([]commonStruct, 0)
		}
		testData[group] = append(testData[group], item)
	}
	for k, _ := range testData {
		t.Run(k, func(t *testing.T) {
			t.Parallel()
			for _, data := range testData[k] {
				size, sizeStr := ParseSize(data.SizeStr)
				if size != data.ExpectSize || sizeStr != data.ExpectSizeStr {
					t.Errorf("测试结果不符合预期: %+v", data)
				}
			}
		})
	}
}

// 模糊测试
func FuzzParseSize(f *testing.F) {
	f.Fuzz(func(t *testing.T, a string) {
		size, sizeStr := ParseSize(a)
		if size == 0 || sizeStr != "" {
			t.Errorf("输入异常导致parseSize没拿到正确结果")
		}
	})
}

// 模糊测试不建议写多个，如果需要，需正则匹配某个特定模糊测试
func FuzzParseSize2(f *testing.F) {
	f.Fuzz(func(t *testing.T, a string) {
		size, sizeStr := ParseSize(a)
		if size == 0 || sizeStr != "" {
			t.Errorf("输入异常导致parseSize没拿到正确结果")
		}
	})
}

// 基准测试(性能测试)
func BenchmarkParseSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseSize("1MB")
	}
}

func BenchmarkParseSizeSub(b *testing.B) {
	testData := make(map[string][]commonStruct)
	for _, item := range commonTestData {
		group := item.Group
		_, ok := testData[group]
		if !ok {
			testData[group] = make([]commonStruct, 0)
		}
		testData[group] = append(testData[group], item)
	}
	b.ResetTimer()
	for k, _ := range testData {
		b.Run(k, func(b *testing.B) {
			//preBenchmark()
			//b.ResetTimer()
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				preBenchmark1()
				b.StartTimer()
				ParseSize(testData[k][0].SizeStr)
			}
		})
	}
}

func preBenchmark() {
	time.Sleep(500 * time.Millisecond)
}
func preBenchmark1() {
	time.Sleep(10000 * time.Nanosecond)
}

// 测试用例的入口函数
func TestMain(m *testing.M) {
	initCommonData()
	// 启动测试
	m.Run()
}

func initCommonData() {
	commonTestData = []commonStruct{
		{"B", "1b", B, "1B"},
		{"B", "100b", 100 * B, "100B"},
		{"KB", "1kb", KB, "1KB"},
		{"KB", "100kb", 100 * KB, "100KB"},
		{"MB", "1Mb", MB, "1MB"},
		//{"MB", "100mB", 100 * MB, "100MB"},
		{"GB", "1Gb", GB, "1GB"},
		{"GB", "10Gb", 10 * GB, "10GB"},
		{"TB", "1tb", TB, "1TB"},
		{"PB", "10PB", 10 * PB, "10PB"},
		{"unknown", "1G", 100 * MB, "100MB"},
	}
}
