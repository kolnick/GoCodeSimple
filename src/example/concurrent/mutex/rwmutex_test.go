package mutex

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRWMutexLock(t *testing.T) {

	var lock sync.RWMutex

	for i := 1; i <= 3; i++ {
		go func(i int) {
			fmt.Printf("go %d 尝试读锁定\n", i)
			lock.RLock()
			fmt.Printf("go %d 已经读锁定了\n", i)
			time.Sleep(5 * time.Second)
			fmt.Printf("go %d 读解锁\n", i)
			lock.RUnlock()
		}(i)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("main try write lock")
	lock.Lock()
	fmt.Println("main write locked")
	lock.Unlock()
	fmt.Println("main write unlock")
}
