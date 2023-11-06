package mutex

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

var tickets int = 20
var wg sync.WaitGroup
var mutex sync.Mutex

// 利用互斥锁实现售票
func TestMutex(t *testing.T) {

	wg.Add(4)

	go saleTickets("1 号窗口", &wg)
	go saleTickets("2 号窗口", &wg)
	go saleTickets("3 号窗口", &wg)
	go saleTickets("4 号窗口", &wg)

	wg.Wait()
	defer fmt.Println("所有车票都售空!")
}

func saleTickets(name string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		mutex.Lock()
		if tickets > 0 {
			time.Sleep(1 * time.Second)
			num, _ := strconv.Atoi(name[:1])
			pre := strings.Repeat("----", num)
			fmt.Println(pre, num, tickets)
			tickets--
		} else {
			fmt.Printf("%s 结束售票\n", name)
			mutex.Unlock()
			break
		}
		mutex.Unlock()
	}

}
