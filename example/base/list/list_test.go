package list

import (
	"container/list"
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	list := list.New()
	list.PushFront(1)
	list.PushFront(2)

	// 遍历输出链表内容
	for e := list.Front(); e != nil; e = e.Next() {
		i := e.Value.(int)
		fmt.Println(i)
	}

	len := list.Len()
	fmt.Println(len)
}
