package list

import (
	"container/list"
	"fmt"
	"testing"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) append(data int) {
	newNode := &Node{data: data, next: nil}

	if ll.head == nil {
		ll.head = newNode
		return
	}

	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (ll *LinkedList) display() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}
func TestSingleLinkedList(t *testing.T) {
	linkedList := &LinkedList{}
	linkedList.append(1)
	linkedList.append(2)
	linkedList.append(3)

	linkedList.display()
}

func TestDoublyLinkedList(t *testing.T) {
	// 创建一个新的双向链表
	myList := list.New()

	// 在列表末尾插入元素
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)

	// 在列表头部插入元素
	myList.PushFront(0)

	// 遍历链表并打印元素
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// 从列表末尾移除一个元素
	myList.Remove(myList.Back())

	// 打印更新后的链表
	fmt.Println("Updated List:")
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
