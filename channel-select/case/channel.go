package _case

import (
	"fmt"
	"strconv"
	"time"
)

// Communication 协程间通讯
func Communication() {
	ch := make(chan int, 10)
	go communicationF1(ch)
	go communicationF2(ch)
}

// 接收一个只写通道
func communicationF1(ch chan<- int) {
	// 通过循环向通道写入数据
	for i := 0; i < 100; i++ {
		ch <- i
	}
}

// 接收一个只读通道
func communicationF2(ch <-chan int) {
	for i := range ch {
		fmt.Println("i = ", i)
	}
}

// ConcurrentSync 并发场景下的同步机制
func ConcurrentSync() {
	// 带缓冲的通道
	ch := make(chan string, 10)
	// 向ch写入消息，写入是无序的
	go func() {
		for i := 0; i < 100; i++ {
			ch <- "协程1写入管道的数据: " + strconv.Itoa(i)
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			ch <- "协程2写入管道的数据: " + strconv.Itoa(i)
		}
	}()
	// 从ch中读取数据并打印
	go func() {
		for val := range ch {
			fmt.Printf("val = %v\n", val)
		}
	}()
}

// NoticeAndMultipleExting 通知协程退出与多路复用
func NoticeAndMultipleExiting() {
	ch := make(chan int)
	strCh := make(chan string)
	done := make(chan struct{})
	go noticeAndMultipleExitingF1(ch)
	go noticeAndMultipleExitingF2(strCh)
	go noticeAndMultipleExitingF3(ch, strCh, done)
	time.Sleep(5 * time.Second)
	close(done)

}

func noticeAndMultipleExitingF1(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
}

func noticeAndMultipleExitingF2(ch chan<- string) {
	for i := 0; i < 100; i++ {
		ch <- fmt.Sprintf("数字: %d", i)
	}
}

// select 子句作为一个整体阻塞，其中任意channel准备好了，就继续执行
func noticeAndMultipleExitingF3(ch <-chan int, strCh <-chan string, done <-chan struct{}) {
	i := 0
label:
	for {
		select {
		case val := <-ch:
			fmt.Println("val = ", val)
		case str := <-strCh:
			fmt.Println("str = ", str)
		case <-done:
			fmt.Println("收到退出通知，退出当前协程")
			break label
			// 没有default语句会阻塞，有default语句会持续执行default子句内容
			// default:
			// fmt.Println("执行default语句")
		}
		i++
	}
	fmt.Println("累计执行次数: ", i)
}
