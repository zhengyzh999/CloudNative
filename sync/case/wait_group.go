package _case

import (
	"fmt"
	"sync"
	"time"
)

// WaitGroup 数据争用的情况，避免使用
func WaitGroup() {
	var a, b = 1000, 10000
	start := time.Now()
	for i := 0; i < 4000000000; i++ {
		multi(a, b)

	}
	step := time.Since(start)
	fmt.Println(step)

	start = time.Now()
	wg := sync.WaitGroup{}
	for i := 0; i < 8; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 500000000; j++ {
				multi(a, b)
			}
		}()
	}
	wg.Wait()
	t := time.Since(start)
	fmt.Println(t)
}

func WaitGroupCase1() {
	ch := make(chan []int, 1000)
	start := time.Now()
	wg2 := sync.WaitGroup{}
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		i := 0
		for item := range ch {
			fmt.Println(multi(item[0], item[1]))
			i++
		}
		time.Sleep(3 * time.Second)
		fmt.Println("数据处理完成，数据条数: ", i)
	}()
	wg := &sync.WaitGroup{}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		wg2.Add(1)
		go func(wg1 *sync.WaitGroup) {
			defer wg1.Done()
			defer wg2.Done()
			for j := 0; j < 500; j++ {
				ch <- []int{j, i}

			}
		}(wg)
	}
	wg.Wait()
	close(ch)
	wg2.Wait()
	t := time.Since(start)
	fmt.Printf("处理时长t = %v\n", t)
}

func multi(a, b int) int {
	return a * b
}
