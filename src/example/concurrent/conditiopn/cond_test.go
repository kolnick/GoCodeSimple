package conditiopn

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Task func()

type ThreadPool struct {
	workers []*Worker
	tasks   chan Task
	cond    *sync.Cond
}

type Worker struct {
	pool *ThreadPool
}

func TestCond(t *testing.T) {
	pool := NewThreadPool(3)
	for i := 0; i < 5; i++ {
		taskId := i
		pool.Submit(func() {
			fmt.Printf("Task %d is being processd\n", taskId)
		})
	}
	// 等待所有任务完成
	pool.cond.L.Lock()

	for len(pool.tasks) > 0 {
		pool.cond.Wait()
	}
	pool.cond.L.Unlock()
}
func (p *ThreadPool) Submit(task Task) {
	p.tasks <- task
}

func (w *Worker) run() {
	for {
		task := <-w.pool.tasks
		task()
	}
}

func NewThreadPool(size int) *ThreadPool {
	pool := &ThreadPool{
		tasks: make(chan Task),
		cond:  sync.NewCond(&sync.Mutex{}),
	}

	for i := 0; i < size; i++ {
		worker := &Worker{pool: pool}
		pool.workers = append(pool.workers, worker)
		go worker.run()
	}
	return pool
}

func TestCond2(t *testing.T) {
	var mutex sync.Mutex
	cond := sync.Cond{L: &mutex}
	condition := false

	go func() {
		time.Sleep(1 * time.Second)
		cond.L.Lock()
		fmt.Println("子 go 已经被锁定")
		fmt.Println("子 go 更改条件,发送通知")
		condition = true
		cond.Signal()
		fmt.Println("子 go 继续")
		time.Sleep(5 * time.Second)
		fmt.Println("子 go 解锁")
		cond.L.Unlock()
	}()

	cond.L.Lock()
	fmt.Println("main 锁定")
	if !condition {
		fmt.Println("main 即将等待")
		cond.Wait()
		fmt.Println("main 被唤醒")
	}
	fmt.Println("main 继续")
	fmt.Println("main 解锁")
	cond.L.Unlock()
}
