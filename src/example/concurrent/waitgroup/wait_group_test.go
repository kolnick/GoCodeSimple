package waitgroup

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"testing"
	"time"
)

// 所有的子go程运行结束后主go才退出
// 类似java的CouldDownLaunch
func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup

	fmt.Printf("%T \n", wg)

	wg.Add(3)

	rand.Seed(time.Now().UnixNano())

	go printNum(&wg, 1)
	go printNum(&wg, 2)
	go printNum(&wg, 3)

	wg.Wait()
	defer fmt.Println("main over...")

}

func printNum(wg *sync.WaitGroup, num int) {
	for i := 1; i <= 3; i++ {
		repeat := strings.Repeat("\t", num-1)
		fmt.Printf("%s 第%d子go 程  %d \n", repeat, num, i)
		time.Sleep(time.Second)
	}
	defer wg.Done()
}
