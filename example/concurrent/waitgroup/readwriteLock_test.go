package waitgroup

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var totalNum int
var wg sync.WaitGroup

// 加入互斥锁 性能效率相对来说是比较低的
var lock sync.RWMutex

func read() {
	defer wg.Done()
	lock.RLock()
	fmt.Println("开始读取数据")
	time.Sleep(time.Second)
	fmt.Println("读取数据成功")
	lock.RUnlock()
}

func write() {
	defer wg.Done()
	lock.Lock()
	fmt.Println("开始修改数据")
	time.Sleep(time.Second * 3)
	fmt.Println("修改数据成功")
	lock.Unlock()
}

func TestReadWriteLock(t *testing.T) {
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go read()
	}
	go write()
	wg.Wait()
}
