package _case

import (
	"fmt"
	"sync"
)

type OnceMap struct {
	sync.Once
	data map[string]int
}

func (m *OnceMap) LoadData() {
	// 内嵌结构体，只要没有重名方法，即可直接使用内嵌结构体的方法
	//m.Once.Do()
	m.Do(func() {
		list := []string{"A", "B", "C", "D", "E", "F"}
		for _, item := range list {
			_, ok := m.data[item]
			if !ok {
				m.data[item] = 0
			}
			m.data[item]++
		}
	})
	// 会出现并发写问题
	/*list := []string{"A", "B", "C", "D", "E", "F"}
	for _, item := range list {
		_, ok := m.data[item]
		if !ok {
			m.data[item] = 0
		}
		m.data[item]++
	}*/
}

func OnceCase() {
	oMap := &OnceMap{
		data: make(map[string]int),
	}
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			oMap.LoadData()
		}()
	}
	wg.Wait()
	fmt.Println(oMap.data)
}
