package _case

import (
	"context"
	"fmt"
	"time"
)

func ContextCase() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "desc", "ContextCase")
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	data := [][]int{{1, 2}, {3, 2}}
	ch := make(chan []int)
	go calculate(ctx, ch)
	for i := 0; i < len(data); i++ {
		ch <- data[i]
	}
	time.Sleep(time.Second * 10)
	// 关闭通道方式通知协程退出
	/*done := make(chan struct{})
	go f1(done)
	go f1(done)
	time.Sleep(time.Second * 3)
	close(done)*/
}

func calculate(ctx context.Context, data <-chan []int) {
	for {
		select {
		case item := <-data:
			ctx = context.WithValue(ctx, "desc", "calculate")
			ch := make(chan []int)
			go sumContext(ctx, ch)
			ch <- item
			ch1 := make(chan []int)
			go multipleContext(ctx, ch1)
			ch1 <- item
		case <-ctx.Done():
			desc := ctx.Value("desc").(string)
			fmt.Println("协程退出, context's desc = ", desc, "。错误消息: ", ctx.Err())
			return
		}
	}
}

func sumContext(ctx context.Context, data <-chan []int) {
	for {
		select {
		case item := <-data:
			a, b := item[0], item[1]
			res := sum(a, b)
			fmt.Printf("%d + %d = %d\n", a, b, res)
		case <-ctx.Done():
			desc := ctx.Value("desc").(string)
			fmt.Println("sumContext协程退出, context's desc = ", desc, "。错误消息: ", ctx.Err())
			return
		}
	}
}

func multipleContext(ctx context.Context, data <-chan []int) {
	for {
		select {
		case item := <-data:
			a, b := item[0], item[1]
			res := multiple(a, b)
			fmt.Printf("%d * %d = %d\n", a, b, res)
		case <-ctx.Done():
			desc := ctx.Value("desc").(string)
			fmt.Println("multipleContext协程退出, context's desc = ", desc, "。错误消息: ", ctx.Err())
			return
		}
	}
}

func sum(a, b int) int {
	return a + b
}

func multiple(a, b int) int {
	time.Sleep(5 * time.Second)
	return a * b
}

func f1(done chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("协程退出")
			return
		}
	}
}
