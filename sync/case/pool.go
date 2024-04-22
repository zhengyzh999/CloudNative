package _case

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
)

func PoolCase() {
	target := "192.168.239.121"
	pool, err := GetPool(target)
	if err != nil {
		log.Fatal(err)
	}
	for i := 1; i <= 5; i++ {
		conn := &Conn{
			Id:     int64(i),
			Target: target,
			Status: ON,
		}
		pool.PutConn(conn)
	}
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 5; i++ {
				conn := pool.GetConn(target)
				pool.PutConn(conn)
				fmt.Println("A", conn.Id)
			}
		}()
	}
	wg.Wait()
}

const (
	ON  = 1
	OFF = 0
)

type Conn struct {
	Id     int64
	Target string
	Status int
}

func (c *Conn) GetStatus() int {
	return c.Status
}
func NewConn(target string) *Conn {
	return &Conn{
		Id:     rand.Int63(),
		Target: target,
		Status: ON,
	}
}

type ConnPool struct {
	sync.Pool
}

func GetPool(target string) (*ConnPool, error) {
	return &ConnPool{
		Pool: sync.Pool{
			New: func() any {
				return NewConn(target)
			},
		},
	}, nil
}

// GetConn get
func (cp *ConnPool) GetConn(target string) *Conn {
	conn := cp.Pool.Get().(*Conn)
	if conn.GetStatus() == OFF {
		conn = cp.Pool.New().(*Conn)
	}
	return conn
}

// PutConn put
func (cp *ConnPool) PutConn(conn *Conn) {
	if conn.GetStatus() == OFF {
		return
	}
	cp.Pool.Put(conn)
}
