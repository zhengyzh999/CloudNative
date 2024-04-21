package _case

import (
	"fmt"
	"sync"
	"time"
)

func CondCase() {
	list := make([]int, 0)
	cond := sync.NewCond(&sync.Mutex{})
	go readList(&list, cond)
	go readList(&list, cond)
	go readList(&list, cond)
	time.Sleep(time.Second * 3)
	initList(&list, cond)
}

func initList(list *[]int, c *sync.Cond) {
	// 主叫方，可以持锁，也可以补持锁
	c.L.Lock()
	defer c.L.Unlock()
	for i := 0; i < 10; i++ {
		*list = append(*list, i)
	}
	// 唤醒所有等待的协程
	c.Broadcast()
	// 唤醒一个等待的协程
	//c.Signal()
}

func readList(list *[]int, c *sync.Cond) {
	// 被叫方，必须持锁
	c.L.Lock()
	defer c.L.Unlock()
	for len(*list) == 0 {
		fmt.Println("readList wait")
		c.Wait()
	}
	fmt.Println("list = ", *list)
}

type Queue struct {
	list []int
	cond *sync.Cond
}

func CondQueueCase() {
	q := newQueue()
	var wg sync.WaitGroup
	for n := 1; n <= 10; n++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			list := q.GetMany(10)
			fmt.Printf("n=%v,list=%v\n", n, list)
		}(n)

	}
	for i := 0; i < 100; i++ {
		q.Put(i)
	}
	wg.Wait()
}

func newQueue() *Queue {
	q := &Queue{
		list: []int{},
		cond: sync.NewCond(&sync.Mutex{}),
	}
	return q
}

func (q *Queue) Put(item int) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.list = append(q.list, item)
	// 通知(唤醒)一个协程来读取
	q.cond.Signal()
}
func (q *Queue) GetMany(n int) []int {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	for len(q.list) < n {
		q.cond.Wait()
	}
	list := q.list[:n]
	q.list = q.list[n:]
	return list
}

/*func CondCase() {
	list := make([]int, 0)
	list = append(list, 1, 2, 3, 4)
	fmt.Printf("main指向地址: %p\n", list)
	fmt.Printf("main变量地址: %p\n", &list)
	fmt.Println("--------------------")
	initList(list)
	fmt.Println("--------------------")
	fmt.Printf("调用方法后指向地址: %p\n", list)
	fmt.Printf("调用方法后变量地址: %p\n", &list)
}*/

// go语言中方法传参全部是值拷贝。即使引用数据类型，也是将变量中保存的地址传递给方法。使得两个方法中的局部变量指向同一个切片，
// 但是两个变量本身的地址不同。如果在方法内部只是对切片数据修改，外部可以得到修改后的值。但是如果使用append操作， 内部方法的
// 局部变量就不会指向原切片，而是新的切片空间(list = append(list,xx))。原切片没有变化，外部变量也就访问不到append后的值
/*func initList(list []int) {
	fmt.Printf("func指向地址: %p\n", list)
	fmt.Printf("func变量地址: %p\n", &list)
	list[2] = 10
	list = append(list, 0)
	fmt.Printf("更改数组后func指向地址: %p\n", list)
	fmt.Printf("更改数组后func变量地址: %p\n", &list)
}*/
