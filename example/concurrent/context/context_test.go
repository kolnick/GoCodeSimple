package context

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s received cancellation signal, exiting...\n", name)
			return
		default:
			fmt.Printf("%s is working...\n", name)
			time.Sleep(1 * time.Second)
		}
	}
}

func TestContext(t *testing.T) {
	// 创建一个父 Context
	parentCtx := context.Background()

	// 派生出两个子 Context
	childCtx1, cancel1 := context.WithCancel(parentCtx)
	childCtx2, cancel2 := context.WithTimeout(parentCtx, 5*time.Second)

	// 启动两个 Goroutine 监听子 Context
	go worker(childCtx1, "Worker 1")
	go worker(childCtx2, "Worker 2")

	// 等待一段时间，然后取消一个子 Context
	time.Sleep(3 * time.Second)
	cancel1() // 取消 Worker 1

	// 等待一段时间，然后查看另一个子 Context 是否超时
	time.Sleep(3 * time.Second)
	cancel2() // 取消 Worker 2

	// 在这里等待一段时间，以观察输出结果
	time.Sleep(2 * time.Second)
}

type keyType string

func TestWithValue(t *testing.T) {
	// 创建一个根 Context
	ctx := context.Background()

	// 使用 WithValue 创建一个派生的 Context，并将键值对关联到其中
	ctxWithValue := context.WithValue(ctx, keyType("userID"), 123)
	var wg sync.WaitGroup
	wg.Add(1)
	// 在 Goroutine 中使用 Context 获取键值对的信息
	go func(ctx context.Context) {
		defer wg.Done()
		userID := ctx.Value(keyType("userID"))
		if userID != nil {
			fmt.Println("User ID:", userID)
		} else {
			fmt.Println("User ID not found")
		}
	}(ctxWithValue)

	// 等待一段时间，以确保 Goroutine 有足够时间执行
	// 如果不等待，程序可能在 Goroutine 执行前就结束了
	wg.Wait()
}
